package domain

import "time"

// TODO: Описать события домена
//  - Событие UserRegistered (для Kafka)
//  - Поля: UserID, Email, RegisteredAt

// UserRegisteredEvent — событие регистрации пользователя
type UserRegisteredEvent struct {
	UserID       UUID
	Email        string
	RegisteredAt time.Time
}
