package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword return hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if error != nil {
		return "", fmt.Errorf("failed to create hashed password: %w", error)
	}

	return string(hashedPassword), nil
}

// CheckPassword compared password with hashedPassword
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
