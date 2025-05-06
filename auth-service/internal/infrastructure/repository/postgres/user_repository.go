package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/domain"
	"time"
)

// TODO: Реализовать методы для работы с пользователями в PostgreSQL

var ErrUserNotFound = errors.New("user not found")

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (id, email, password_hash, created_at)
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email.String(), user.PasswordHash, user.CreatedAt)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	query := `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`

	var u struct {
		ID           domain.UUID `db:"id"`
		Email        string      `db:"email"`
		PasswordHash string      `db:"password_hash"`
		CreatedAt    time.Time   `db:"created_at"`
	}

	err := r.db.GetContext(ctx, &u, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, fmt.Errorf("get user by email: %w", err)
	}

	emailObj, err := domain.NewEmail(u.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("invalid email in db: %w", err)
	}

	return domain.User{
		ID:           u.ID,
		Email:        emailObj,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
	}, nil
}
