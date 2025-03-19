package main

import (
	"fmt"
)

type TagManager struct {
	List map[string]struct{}
}

func NewTagManager() *TagManager {
	return &TagManager{
		List: make(map[string]struct{}),
	}
}

func (t *TagManager) AddTag(tag string) error {
	if t.TagExists(tag) {
		return fmt.Errorf("tag %s already exists", tag)
	}
	t.List[tag] = struct{}{}
	return nil
}

func (t TagManager) TagExists(tag string) bool {
	_, ok := t.List[tag]
	return ok
}

func (t TagManager) ListTags() (result []string) {
	for key := range t.List {
		result = append(result, key)
	}
	return
}

func (t *TagManager) RemoveTag(tag string) error {
	if !t.TagExists(tag) {
		return fmt.Errorf("tag %s not found", tag)
	}
	delete(t.List, tag)
	return nil
}

func main() {
	tm := NewTagManager()

	// Добавление тегов
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("programming"); err != nil {
		fmt.Println(err)
	}
	if err := tm.AddTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег уже существует
	}

	// // Проверка существования тегов
	fmt.Println("Тег 'golang' существует:", tm.TagExists("golang")) // true
	fmt.Println("Тег 'python' существует:", tm.TagExists("python")) // false

	// Список тегов
	fmt.Println("Current tags:", tm.ListTags()) // [golang programming]

	// Удаление тегов
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err)
	}
	if err := tm.RemoveTag("golang"); err != nil {
		fmt.Println(err) // Ошибка, тег не существует
	}

	// // Список тегов после удаления
	fmt.Println("Current tags after removal:", tm.ListTags()) // [programming]
}
