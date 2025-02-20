package main

import (
	"fmt"
	"slices"
	"strings"
)

func countFriends(m map[string][]string) map[string]int {
	result := make(map[string]int)

	for name, friendsName := range m {
		i := 0
		for range friendsName {
			i++
		}
		result[name] = i
	}
	return result
}

func commonFriends(m map[string][]string, name1, name2 string) []string {
	result := make([]string, 0)
	for _, friendsName := range m[name1] {
		if slices.Contains(m[name2], friendsName) {
			result = append(result, friendsName)
		}
	}
	return result
}

func mostPopularUsers(m map[string][]string) ([]string, int) {
	countFriendsList := make([]int, 0)
	for _, count := range countFriends(m) {
		countFriendsList = append(countFriendsList, count)
	}
	maxCount := slices.Max(countFriendsList)

	result := make([]string, 0)
	friendAndCount := countFriends(m)
	for friendsName, count := range friendAndCount {
		if count == maxCount {
			result = append(result, friendsName)
		}
	}
	return result, maxCount
}

func main() {
	friendsData := map[string][]string{
		"Алексей":  {"Иван", "Сергей", "Елена"},
		"Иван":     {"Алексей", "Дмитрий", "Мария"},  //
		"Елена":    {"Алексей", "Сергей", "Дмитрий"}, //
		"Сергей":   {"Алексей", "Елена"},
		"Дмитрий":  {"Иван", "Елена", "Ольга"},
		"Мария":    {"Иван", "Ольга"},
		"Ольга":    {"Дмитрий", "Мария"},
		"Анна":     {"Петр"},
		"Петр":     {"Анна", "Сергей"},
		"Светлана": {"Иван", "Елена"},
	}
	nameAndCountCountFriends := countFriends(friendsData)
	fmt.Println("Количество друзей:")
	listCountFriends := make([]string, 0, len(nameAndCountCountFriends))
	for name := range nameAndCountCountFriends {
		listCountFriends = append(listCountFriends, name)
	}
	slices.Sort(listCountFriends)
	//вывод:
	for _, name := range listCountFriends {
		fmt.Printf("%s: %d\n", name, nameAndCountCountFriends[name])
	}

	nameCommonFriends := commonFriends(friendsData, "Иван", "Елена")
	slices.Sort(nameCommonFriends)
	fmt.Printf("Общие друзья между пользователями Иван и Елена: %s.\n", strings.Join(nameCommonFriends, ", "))

	listNameCommonFriends, maxCountCommonFriends := mostPopularUsers(friendsData)
	slices.Sort(listNameCommonFriends)
	fmt.Printf("Наиболее популярные пользователи: %s (количество друзей: %d).\n", strings.Join(listNameCommonFriends, ", "), maxCountCommonFriends)
}
