package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/command"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/interfaces"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/query"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/infrastructure/repository/postgres"
	"time"
)

type Service struct {
	userRepository interfaces.UserRepository
	sessionStore   interfaces.SessionStore
	tokenManager   interfaces.TokenManager
	passwordHash   domain.PasswordHash
	uuidGenerator  domain.UUIDGenerator
	clock          domain.Clock
}

func NewService(
	userRepository interfaces.UserRepository,
	sessionStore interfaces.SessionStore,
	tokenManager interfaces.TokenManager,
	hash domain.PasswordHash,
	uuidGen domain.UUIDGenerator,
	clock domain.Clock,
) *Service {
	return &Service{
		userRepository: userRepository,
		sessionStore:   sessionStore,
		tokenManager:   tokenManager,
		passwordHash:   hash,
		uuidGenerator:  uuidGen,
		clock:          clock,
	}
}

func (s *Service) Register(ctx context.Context, cmd command.RegisterCommand) error {
	// Проверяем, не существует ли пользователь с таким email
	_, err := s.userRepository.GetUserByEmail(ctx, cmd.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", cmd.Email)
	}
	if !errors.Is(err, postgres.ErrUserNotFound) {
		return fmt.Errorf("could not check user existence: %w", err)
	}

	newUser, err := domain.NewUserFromRegister(
		s.uuidGenerator,
		cmd.Email,
		cmd.Password,
		s.clock,
		s.passwordHash,
	)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("could not save user: %w", err)
	}

	return nil
}

func (s *Service) Login(ctx context.Context, cmd query.LoginQuery) (domain.TokenPair, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not get user: %w", err)
	}

	password, err := domain.NewPassword(cmd.Password)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not parse password: %w", err)
	}

	err = user.CheckPassword(password, s.passwordHash)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("invalid password: %w", err)
	}

	tokenAccess, err := s.tokenManager.GenerateAccessToken(user.ID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not generate access token: %w", err)
	}

	tokenRefresh, err := s.tokenManager.GenerateRefreshToken(user.ID)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not generate refresh token: %w", err)
	}

	err = s.sessionStore.SaveRefreshToken(ctx, user.ID, tokenRefresh, time.Hour*24*30)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not save refresh token: %w", err)
	}

	return domain.TokenPair{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (domain.TokenPair, error) {
	userIDFromJWT, err := s.tokenManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not validate refresh token: %w", err)
	}

	userIDFromRedis, err := s.tokenManager.GetUserIDByRefreshToken(ctx, refreshToken)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not get user ID by refresh token: %w", err)
	}

	// Проверяем, что токен не был подделан (ID из Redis и JWT должны совпадать)
	err = validateUserIDConsistency(userIDFromJWT, userIDFromRedis)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not validate refresh token: %w", err)
	}

	newAccessToken, err := s.tokenManager.GenerateAccessToken(userIDFromJWT)
	if err != nil {
		return domain.TokenPair{}, fmt.Errorf("could not generate access token: %w", err)
	}

	newRefreshToken, err := s.generateAndStoreRefreshToken(ctx, userIDFromJWT)
	if err != nil {
		return domain.TokenPair{}, err
	}

	return domain.TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

// validateUserIDConsistency проверяет, совпадают ли ID из JWT и из Redis.
func validateUserIDConsistency(userIDFromJWT, userIDFromRedis uuid.UUID) error {
	if userIDFromRedis != userIDFromJWT {
		return fmt.Errorf("invalid userID from redis")
	}
	return nil
}

// generateAndStoreRefreshToken генерирует новый refresh токен и сохраняет его в sessionStore.
func (s *Service) generateAndStoreRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	refreshToken, err := s.tokenManager.GenerateRefreshToken(userID)
	if err != nil {
		return "", fmt.Errorf("could not generate refresh token: %w", err)
	}

	err = s.sessionStore.SaveRefreshToken(ctx, userID, refreshToken, time.Hour*24*30)
	if err != nil {
		return "", fmt.Errorf("could not save refresh token: %w", err)
	}

	return refreshToken, nil
}
