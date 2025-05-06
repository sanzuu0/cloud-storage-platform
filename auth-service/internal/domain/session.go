package domain

import "time"

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
