package adapters

// TODO: Адаптер генерации UUID
//  - Генерировать UUID
//  - Нужен для мокирования в тестах

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func (GoogleUUIDGenerator) NewUUID() uuid.UUID {
	return uuid.New()
}
