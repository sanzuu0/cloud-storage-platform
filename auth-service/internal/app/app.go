package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/config"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/command"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/infrastructure/adapters"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/infrastructure/repository/postgres"
	"log"
	"time"
)

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

type stubSessionStore struct{}
type stubTokenManager struct{}

func Run(cfg config.Config) error {

	db, err := sqlx.Connect("postgres", cfg.PostgresDSN)
	if err != nil {
		return err
	}

	userRepo := postgres.NewUserRepository(db)
	sessionStore := stubSessionStore{}
	tokenManager := stubTokenManager{}
	passwordHash := adapters.BcryptHash{}
	uuidGenerator := adapters.GoogleUUIDGenerator{}
	clock := adapters.SystemClock{}

	// Сборка сервиса
	service := application.NewService(
		userRepo,
		sessionStore,
		tokenManager,
		passwordHash,
		uuidGenerator,
		clock,
	)

	err = service.Register(context.Background(), command.RegisterCommand{
		Email:    "test@example.com",
		Password: "P@ssw0rd123!",
	})
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}
	log.Println("✅ User successfully registered!")

	log.Printf("Service initialized: %+v", service)

	return nil
}
