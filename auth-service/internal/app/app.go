package app

import (
	"context"
	"github.com/google/uuid"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/config"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/infrastructure/adapters"
	"log"
	"time"
)

// --- заглушка UserRepository ---
func (stubUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	return nil
}
func (stubUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

// --- заглушка SessionStore ---
func (stubSessionStore) Save(ctx context.Context, session domain.Session) error {
	return nil
}
func (stubSessionStore) SaveRefreshToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error {
	return nil
}
func (stubSessionStore) SaveAccessToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error {
	return nil
}
func (stubSessionStore) GetUserID(ctx context.Context, refreshToken string) (domain.UUID, error) {
	return uuid.New(), nil
}
func (stubSessionStore) Delete(ctx context.Context, refreshToken string) error {
	return nil
}

// --- заглушка TokenManager ---
func (stubTokenManager) GenerateAccessToken(userID domain.UUID) (string, error) {
	return "access-token", nil
}
func (stubTokenManager) GenerateRefreshToken(userID domain.UUID) (string, error) {
	return "refresh-token", nil
}
func (stubTokenManager) ValidateAccessToken(token string) (domain.UUID, error) {
	return uuid.New(), nil
}
func (stubTokenManager) ValidateRefreshToken(token string) (domain.UUID, error) {
	return uuid.New(), nil
}
func (stubTokenManager) GetUserIDByRefreshToken(ctx context.Context, token string) (uuid.UUID, error) {
	return uuid.New(), nil
}

type stubUserRepository struct{}
type stubSessionStore struct{}
type stubTokenManager struct{}

func Run(cfg config.Config) error {

	passwordHash := adapters.NewBcryptHash()
	uuidGenerator := adapters.NewGoogleUUIDGenerator()
	clock := adapters.NewSystemClock()

	// TODO: временные заглушки для репозиториев
	userRepo := stubUserRepository{}
	sessionStore := stubSessionStore{}
	tokenManager := stubTokenManager{}

	// Сборка сервиса
	service := application.NewService(
		userRepo,
		sessionStore,
		tokenManager,
		passwordHash,
		uuidGenerator,
		clock,
	)

	// Пока просто логируем успешную инициализацию
	log.Printf("Service initialized: %+v", service)

	// TODO: Здесь будет запуск сервера (HTTP или gRPC)

	return nil
}
