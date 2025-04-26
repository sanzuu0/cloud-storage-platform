package domain

import "errors"

// TODO: Описать бизнес-ошибки
//  - UserAlreadyExists
//  - InvalidCredentials
//  - SessionNotFound

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrSessionNotFound    = errors.New("session not found")
)
