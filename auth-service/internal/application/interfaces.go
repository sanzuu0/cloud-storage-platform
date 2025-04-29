package application

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type SessionStore interface {
	Save(ctx context.Context, session domain.Session) error
	SaveRefreshToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error
	SaveAccessToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error
	GetUserID(ctx context.Context, refreshToken string) (domain.UUID, error)
	Delete(ctx context.Context, refreshToken string) error
}

type TokenManager interface {
	GenerateAccessToken(userID domain.UUID) (string, error)
	GenerateRefreshToken(userID domain.UUID) (string, error)
	ValidateAccessToken(token string) (domain.UUID, error)
	ValidateRefreshToken(token string) (domain.UUID, error)
}
