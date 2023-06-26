package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while hashing password: %w", err)
	}

	return string(hash), nil
}

func comparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
