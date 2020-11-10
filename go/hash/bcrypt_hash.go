package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(plain string) (string, error) {
	b := []byte(plain)
	b, err := bcrypt.GenerateFromPassword(b, 10)

	return string(b), nil
}

func Check(plain, hashed string) error {
	p := []byte(plain)
	h := []byte(hashed)

	return bcrypt.CompareHashAndPassword(h, p)
}
