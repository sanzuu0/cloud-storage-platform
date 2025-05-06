package adapters

import (
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func NewGoogleUUIDGenerator() *GoogleUUIDGenerator {
	return &GoogleUUIDGenerator{}
}

func (GoogleUUIDGenerator) NewUUID() uuid.UUID {
	return uuid.New()
}
