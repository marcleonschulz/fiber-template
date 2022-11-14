package hash

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error(err)
	}
	if CheckPasswordHash(password, hash) == false {
		t.Error("Hashing failed")
	}
}
