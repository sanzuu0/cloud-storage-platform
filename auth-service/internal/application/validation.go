package application

import (
	"fmt"
	"github.com/sanzuu0/cloud-storage-platform/auth-service/internal/application/command"
	"regexp"
)

// проверка формата почты (___@__.__)
func emailValidator(cmd command.RegisterCommand) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(cmd.Email) {
		err := fmt.Errorf("invalid email format")
		return err
	}
	return nil
}

// проверка пароля на символы, спец символы и длину.
func passwordValidator(cmd command.RegisterCommand) error {
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(cmd.Password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(cmd.Password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(cmd.Password)

	if len(cmd.Password) < 8 || !hasLetter || !hasNumber || !hasSpecial {
		err := fmt.Errorf("password must contain at least 8 characters")
		return err
	}
	return nil
}
