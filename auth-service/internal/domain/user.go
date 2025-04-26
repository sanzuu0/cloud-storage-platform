package domain

import (
	"time"
)

// TODO: Описать сущность пользователя
//  - UUID идентификатор
//  - Email
//  - PasswordHash
//  - CreatedAt

// TODO: Возможно добавить методы валидации (например, проверка корректности Email)

// USER - сущность пользователя в сисеме

type User struct {
	ID           UUID
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(id UUID, email string, password string, createdAt time.Time) User {
	return User{
		ID:           id,
		Email:        email,
		PasswordHash: password,
		CreatedAt:    createdAt,
	}
}
