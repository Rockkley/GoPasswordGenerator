package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"strings"
)

const (
	MIN_PASSWORD_LENGTH = 8
	MAX_PASSWORD_LENGTH = 16
	PASSWORD_CHARSET    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-="
)

func generatePassword(length int) (string, error) {
	var generatedPassword string
	if length < MIN_PASSWORD_LENGTH || length > MAX_PASSWORD_LENGTH {
		fmt.Println("Invalid password length")
		return "", fmt.Errorf("")
	}
	var stringBuilder strings.Builder
	stringBuilder.Grow(length)
	byteArray := make([]byte, length)
	_, err := cryptoRand.Read(byteArray)
	if err != nil {
		return "", err
	}
	for i := 0; i < length; i++ {
		stringBuilder.WriteString(string(PASSWORD_CHARSET[int(byteArray[i])%len(PASSWORD_CHARSET)]))
	}
	generatedPassword = stringBuilder.String()
	return generatedPassword, nil
}
func main() {
	var length = MIN_PASSWORD_LENGTH
	fmt.Printf("Enter length of password (minimum %d, maximum %d characters): ",
		MIN_PASSWORD_LENGTH, MAX_PASSWORD_LENGTH)
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Println(err)
		return
	}

	password, _ := generatePassword(length)
	fmt.Println(password)
}
