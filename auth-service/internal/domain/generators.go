package domain

import (
	"github.com/google/uuid"
	"time"
)

type UUIDGenerator interface {
	NewUUID() uuid.UUID
}

type Clock interface {
	Now() time.Time
}
