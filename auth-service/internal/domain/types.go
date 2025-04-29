package domain

import (
	"github.com/google/uuid"
)

// TODO: Описать базовые типы и вспомогательные функции
//  - UUID тип
//  - Email тип
//  - Вспомогательные методы: создание UUID, валидация Email

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type UUID = uuid.UUID
