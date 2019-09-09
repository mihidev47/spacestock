package crypto

import (
	"golang.org/x/crypto/bcrypt"

	"../logger"
)

const (
	costFactor = 14
)

var log = logger.Get()

// HashPassword generates hash string from password input
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return string(bytes), err
}

// CheckPassword check input password against hash
func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
