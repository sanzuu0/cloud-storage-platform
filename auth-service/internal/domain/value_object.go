package domain

import (
	"fmt"
	"regexp"
)

// Email — value object, проверяет и инкапсулирует email
type Email string

// Password — value object, валидирует и инкапсулирует сырой пароль
type Password string

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	hasLetter  = regexp.MustCompile(`[a-zA-Z]`)
	hasNumber  = regexp.MustCompile(`[0-9]`)
	hasSpecial = regexp.MustCompile(`[!@#\$%\^&\*]`)
)

func (e Email) String() string {
	return string(e)
}

func NewEmail(line string) (Email, error) {
	if !emailRegex.MatchString(line) {
		err := fmt.Errorf("invalid email format")
		return "", err
	}
	return Email(line), nil
}

func (p Password) String() string {
	return string(p)
}

func NewPassword(line string) (Password, error) {

	if len(line) < 8 || !hasLetter.MatchString(line) || !hasNumber.MatchString(line) || !hasSpecial.MatchString(line) {
		err := fmt.Errorf("password must be at least 8 characters and include letters, numbers, and special characters")
		return "", err
	}
	return Password(line), nil
}
