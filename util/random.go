package util

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomTestStrin() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@whatever.com", RandomString(4))
}

func RandomPassword() (string, error) {
	pass := RandomString(6)

	hashedPassword, err := HashPass(pass)
	if err != nil {
		log.Fatal("error generating random password")
		return "", err
	}

	return hashedPassword, nil
}
