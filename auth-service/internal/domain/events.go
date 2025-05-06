package domain

import "time"

type UserRegisteredEvent struct {
	UserID       UUID
	Email        string
	RegisteredAt time.Time
}
