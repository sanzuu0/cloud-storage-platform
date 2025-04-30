package interfaces

import (
	"context"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}
