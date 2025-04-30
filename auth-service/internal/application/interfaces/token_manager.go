package interfaces

import (
	"context"
	"github.com/google/uuid"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
)

type TokenManager interface {
	GenerateAccessToken(userID domain.UUID) (string, error)
	GenerateRefreshToken(userID domain.UUID) (string, error)
	ValidateAccessToken(token string) (domain.UUID, error)
	ValidateRefreshToken(token string) (domain.UUID, error)
	GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error)
}
