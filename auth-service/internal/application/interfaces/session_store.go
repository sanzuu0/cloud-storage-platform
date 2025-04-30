package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"time"
)

type SessionStore interface {
	Save(ctx context.Context, session domain.Session) error
	SaveRefreshToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error
	SaveAccessToken(ctx context.Context, uuid uuid.UUID, token string, duration time.Duration) error
	GetUserID(ctx context.Context, refreshToken string) (domain.UUID, error)
	Delete(ctx context.Context, refreshToken string) error
}
