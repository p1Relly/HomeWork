package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 4, 5}
	numbers2 := []int{1, 2, 4, 5, -100}
	newnumbers := slices.Insert(numbers, 2, 3) // вставить значение в определенный элемент
	fmt.Println(numbers, newnumbers)

	fmt.Println("поиск значения в слайсе:", slices.Contains(numbers, 4))                      // поиск значения в слайсе
	fmt.Println("поиск по функции в слайсе:", slices.ContainsFunc(numbers, func(v int) bool { // поиск значений по функции (четные)
		return v%2 == 0
	}))

	fmt.Println("сравнение слайса 1 и слайса 2:", slices.Compare(numbers, numbers2)) //сравнение 1 - больше, -1 - меньше, 0 - равны

	slices.CompareFunc()
}
