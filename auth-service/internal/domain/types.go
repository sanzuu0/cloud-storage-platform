package domain

import (
	"errors"
	"github.com/google/uuid"
	"regexp"
)

// TODO: Описать базовые типы и вспомогательные функции
//  - UUID тип
//  - Email тип
//  - Вспомогательные методы: создание UUID, валидация Email

type UUID = uuid.UUID

// isEmailValid проверяет корректность email.
func isEmailValid(email string) error {

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}
