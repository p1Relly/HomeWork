package main

import (
	"errors"
	"fmt"
)

// Данные
type Message struct {
	From    string
	Message string
}

type Chat struct {
	ID       int
	Messages []Message
}

// Ошибка базы данных
type DatabaseError struct {
	Message string
	Code    int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error with code %d: %s", e.Code, e.Message)
}

// Сервис для работы
type Service struct {
	db interface {
		GetChatByIDWithMessages(id int) (*Chat, error)
	}
}

func (w Service) PrintChat(id int) {
	result, err := w.db.GetChatByIDWithMessages(id)
	var ErrDatabase *DatabaseError
	if errors.As(err, &ErrDatabase) && ErrDatabase.Code == 24 {
		fmt.Printf("Ошибка запроса: %s", ErrDatabase.Error())
		return
	}
	if errors.As(err, &ErrDatabase) && ErrDatabase.Code != 24 {
		fmt.Printf("Инфраструктурная ошибка: %s", ErrDatabase.Error())
		return
	}
	if err != nil && !errors.As(err, &ErrDatabase) {
		fmt.Printf("Неизвестная ошибка: %s", err)
		return
	}

	for _, v := range result.Messages {
		fmt.Printf("%s: %s\n", v.From, v.Message)
	}
}

func main() {
}
