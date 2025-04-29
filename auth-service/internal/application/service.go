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

// TODO: Реализовать бизнес-логику аутентификации
//  - Метод RegisterUser(ctx, email, password) (создание пользователя, хеширование пароля, запись в БД, публикация события в Kafka)
//  - Метод Login(ctx, email, password) (проверка пароля, выдача токенов)
//  - Метод RefreshTokens(ctx, refreshToken) (обновление токенов)
//  - Использовать интерфейсы Repository, TokenManager, SessionStore

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
	errEmailValidator := emailValidator(cmd)
	if errEmailValidator != nil {
		return fmt.Errorf("invalid email format: %w", errEmailValidator)
	}

	// проверка формата пароля
	errPasswordValidator := passwordValidator(cmd)
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
		panic(err)
	}

	// создание пользователя
	UUID, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	newUser := domain.NewUser(UUID, cmd.Email, hashPassword, time.Now())

	// запись юзера в бд
	err = s.userRepository.CreateUser(ctx, newUser)

	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (s *Service) Login(ctx context.Context, cmd query.LoginQuery) error {
	return nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) error {
	return nil
}
