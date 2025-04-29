package application

import (
	"fmt"
	"regexp"
)

// проверка формата почты (___@__.__)
func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		err := fmt.Errorf("invalid email format")
		return err
	}
	return nil
}

// проверка пароля на символы, спец символы и длину.
func ValidatePassword(password string) error {
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password)

	if len(password) < 8 || !hasLetter || !hasNumber || !hasSpecial {
		err := fmt.Errorf("password must contain at least 8 characters")
		return err
	}
	return nil
}
