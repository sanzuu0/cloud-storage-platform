package domain

import "time"

// TODO: Описать сущность сессии пользователя (Refresh Token)
//  - ID токена
//  - ID пользователя (UserID)
//  - Время истечения токена (ExpiresAt)

type Session struct {
	Token     string
	UserID    UUID
	ExpiresAt time.Time
}

func NewSession(token string, userID UUID, expiresAt time.Time) *Session {
	return &Session{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}
