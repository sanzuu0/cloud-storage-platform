package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/command"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/query"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/infrastructure/repository/postgres"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	userRepository UserRepository
	sessionStore   SessionStore
	tokenManager   TokenManager
}

func NewService(userRepository UserRepository, sessionStore SessionStore, tokenManager TokenManager) *Service {
	return &Service{
		userRepository: userRepository,
		sessionStore:   sessionStore,
		tokenManager:   tokenManager,
	}
}

func (s *Service) Register(ctx context.Context, cmd command.RegisterCommand) error {

	// проверка формата почты
	errEmailValidator := ValidateEmail(cmd.Email)
	if errEmailValidator != nil {
		return fmt.Errorf("invalid email format: %w", errEmailValidator)
	}

	// проверка формата пароля
	errPasswordValidator := ValidatePassword(cmd.Password)
	if errPasswordValidator != nil {
		return fmt.Errorf("invalid password: %w", errPasswordValidator)
	}

	// существует ли такой пользователь в бд
	_, err := s.userRepository.GetUserByEmail(ctx, cmd.Email)

	if err == nil {
		return fmt.Errorf("user with email %s already exists", cmd.Email)
	}
	if !errors.Is(err, postgres.ErrUserNotFound) {
		return fmt.Errorf("could not check user existence: %w", err)
	}

	// хеширование пароля
	hashBytePassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	hashPassword := string(hashBytePassword)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	// создание пользователя
	UUID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("could not create UUID: %w", err)
	}
	newUser := domain.NewUser(UUID, cmd.Email, hashPassword, time.Now())

	// запись юзера в бд
	err = s.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (s *Service) Login(ctx context.Context, cmd query.LoginQuery) (domain.TokenPair, error) {

	// возьмем пользователя из бд
	user, err := s.userRepository.GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not get user: %w", err)
	}

	// проверка на идентичность паролей
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password))
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("invalid password: %w", err)
	}

	// Генерируем access токен
	tokenAccess, err := s.tokenManager.GenerateAccessToken(user.ID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not generate access token: %w", err)
	}

	// Генерируем refresh токен
	tokenRefresh, err := s.tokenManager.GenerateRefreshToken(user.ID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not generate refresh token: %w", err)
	}

	// сохраняем refresh токен в redis
	err = s.sessionStore.SaveRefreshToken(ctx, user.ID, tokenRefresh, time.Hour*24*30)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not save refresh token: %w", err)
	}

	return domain.TokenPair{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) error {
	return nil
}
