package domain

import (
	"fmt"
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
	Email        Email
	PasswordHash string
	CreatedAt    time.Time
}

func NewUserFromRegister(
	uuid UUIDGenerator,
	rawEmail string,
	rawPassword string,
	clock Clock,
	hashing PasswordHash,
) (User, error) {

	id := uuid.NewUUID()

	email, err := NewEmail(rawEmail)
	if err != nil {
		return User{}, fmt.Errorf("cannot create email: %w", err)
	}

	password, err := NewPassword(rawPassword)
	if err != nil {
		return User{}, fmt.Errorf("error creating password: %w", err)
	}

	hashPassword, err := hashing.Hash(password.String())
	if err != nil {
		return User{}, fmt.Errorf("error hashing password: %w", err)
	}

	timeCreated := clock.Now()

	return User{
		ID:           id,
		Email:        email,
		PasswordHash: hashPassword,
		CreatedAt:    timeCreated,
	}, nil
}

func (u *User) CheckPassword(password Password, hashing PasswordHash) error {
	return hashing.Compare(u.PasswordHash, password.String())
}
