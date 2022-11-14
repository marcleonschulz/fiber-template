package hash

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword takes a password and hash it
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if CheckPasswordHash(password, string(bytes)) == false {
		log.Fatal("Hashing failed")
	}
	return string(bytes), err
}

// CheckPasswordHash takes hashed password and a password and compare them
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
