package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
	"time"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

func (u User) SecretIdentity() string {
	return string([]rune(u.FirstName)[0]) + string([]rune(u.LastName)[0]) + strconv.Itoa(rand.IntN(100)+1)
}

func (u User) Age() int {
	return time.Now().Year() - u.BirthYear
}

func (u *User) AddFavoriteLanguage(language string) error {
	if slices.Contains(u.FavoriteLanguages, language) {
		return errors.New("duplicate")
	}
	if language == "" {
		return errors.New("empty language name")
	}
	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil
}

func (u *User) RemoveFavoriteLanguage(language string) error {
	for i := range u.FavoriteLanguages {
		if u.FavoriteLanguages[i] == language {
			u.FavoriteLanguages = append(u.FavoriteLanguages[:i], u.FavoriteLanguages[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func (u User) IsProgrammingLanguageFavorite(language string) bool {
	return slices.Contains(u.FavoriteLanguages, language)
}

func (u User) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) == 0 {
		return "", errors.New("no options")
	}
	return u.FavoriteLanguages[rand.IntN(len(u.FavoriteLanguages))], nil
}

func (u User) GenerateProfile() string {
	return fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: [%s].", u.FirstName, u.LastName, u.Age(), strings.Join(u.FavoriteLanguages, ", "))
}

func (u *User) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("empty data")
	}
	charFirstName := []rune(firstName)
	charLastName := []rune(lastName)
	if string(charFirstName[0]) == strings.ToLower(string(charFirstName[0])) || string(charLastName[0]) == strings.ToLower(string(charLastName[0])) {
		return errors.New("invalid data")
	}
	u.FirstName = firstName
	u.LastName = lastName
	return nil
}

func main() {
	user1 := User{
		FirstName:         "Иван",
		LastName:          "Владимиров",
		BirthYear:         1996,
		FavoriteLanguages: []string{"go", "c++"},
	}
	fmt.Println("SecretIdentity:", user1.SecretIdentity())
	fmt.Println("Age:", user1.Age())

	AddFavoriteLanguage := user1.AddFavoriteLanguage("SQL")
	if AddFavoriteLanguage != nil {
		fmt.Printf("error: %v\n", AddFavoriteLanguage)
	}
	fmt.Println("AddFavoriteLanguage:", user1.FavoriteLanguages)

	RemoveFavoriteLanguage := user1.RemoveFavoriteLanguage("SQL")
	if AddFavoriteLanguage != nil {
		fmt.Printf("error: %v\n", RemoveFavoriteLanguage)
	}
	fmt.Println("[after] RemoveFavoriteLanguage:", user1.FavoriteLanguages)

	fmt.Println("IsProgrammingLanguageFavorite:", user1.IsProgrammingLanguageFavorite("go"))

	randomFavoriteLanguage, err := user1.RandomFavoriteLanguage()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println("RandomFavoriteLanguage:", randomFavoriteLanguage)
	fmt.Printf("\n%s", user1.GenerateProfile())

	updateName := user1.UpdateName("Вася", "Пупкин")
	if updateName != nil {
		fmt.Printf("error: %v\n", updateName)
	}

	fmt.Printf("\n%s", user1.GenerateProfile())
}
