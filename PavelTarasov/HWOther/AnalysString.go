package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func GetInput() (input string, err error) {
	fmt.Printf("Введите текст: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка ввода: %w", err)
	} else if scanner.Text() == "" {
		return "", fmt.Errorf("Пустой ввод")
	}

	return scanner.Text(), nil
}

func CountCharacters(text string) (letters int, digits int, spaces int, punctuation int) {
	for _, char := range text {
		if unicode.IsDigit(char) {
			digits++
		} else if unicode.IsLetter(char) {
			letters++
		} else if unicode.IsPunct(char) {
			punctuation++
		}
	}

	return letters, digits, strings.Count(text, " "), punctuation
}

func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Количество букв: %d\nКоличество цифр: %d\nКоличество пробелов: %d\nКоличество знаков препинания: %d", letters, digits, spaces, punctuation)
}

func main() {
	input, err := GetInput()
	if err != nil {
		log.Fatalf("Ошибка: %s", err)
	}

	DisplayResults(CountCharacters(input))
}
