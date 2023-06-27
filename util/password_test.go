package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPass(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	// Success case
	err = ComparePassword(password, hashedPassword)
	require.NoError(t, err)

	// Fail case
	wrongPassword := RandomString(4)
	err = ComparePassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
