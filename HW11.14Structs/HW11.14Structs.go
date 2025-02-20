package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrTooLowPasswordsCount = errors.New("too low passwords count")
	ErrTooBigPasswordsCount = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

func getRandomChar(correctChar []rune, m map[string]struct{}) string {
	for {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(correctChar))))
		if _, ok := m[string(correctChar[int(random.Int64())])]; !ok {
			return string(correctChar[int(random.Int64())])
		}
	}
}

func generatePassword(length, count int) ([]string, error) {
	if length < MinPasswordLength {
		return []string{""}, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return []string{""}, ErrTooLowPasswordsCount
	}
	if count > MaxPasswordsCount {
		return []string{""}, ErrTooBigPasswordsCount
	}

	result := make([]string, 0, count)

	for j := 0; j < count; j++ {
		m := make(map[string]struct{}, length)

		m[getRandomChar(upperChars, m)] = struct{}{}
		m[getRandomChar(lowerChars, m)] = struct{}{}
		m[getRandomChar(digitChars, m)] = struct{}{}
		m[getRandomChar(specialChars, m)] = struct{}{}

		for i := 0; i < length-4; i++ {
			random, _ := rand.Int(rand.Reader, big.NewInt(4))
			switch {
			case int(random.Int64()) == 0:
				m[getRandomChar(upperChars, m)] = struct{}{}
			case int(random.Int64()) == 1:
				m[getRandomChar(lowerChars, m)] = struct{}{}
			case int(random.Int64()) == 2:
				m[getRandomChar(digitChars, m)] = struct{}{}
			case int(random.Int64()) == 3:
				m[getRandomChar(specialChars, m)] = struct{}{}
			}
		}

		val := ""
		for key := range m {
			val += key
		}
		result = append(result, val)
	}

	return result, nil
}

func main() {
	passwords, err := generatePassword(12, 4)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(passwords) // Вывод сгенерированных паролей
}
