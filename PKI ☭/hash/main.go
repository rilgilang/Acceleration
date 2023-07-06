package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Hash password using the bcrypt hashing algorithm
func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func main() {
	// Hash password
	var hashedPassword, err = hashPassword("GanyangPKI☭")

	if err != nil {
		println(fmt.Println("Error hashing password"))
		return
	}

	fmt.Println("Password Hash:", hashedPassword)
}
