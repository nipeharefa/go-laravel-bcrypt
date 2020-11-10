package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	b := []byte("nipeharefa")
	b, err := bcrypt.GenerateFromPassword(b, 10)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
