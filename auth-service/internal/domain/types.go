package domain

import "github.com/google/uuid"

type UUID = uuid.UUID

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}
