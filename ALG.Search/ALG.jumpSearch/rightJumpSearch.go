package main

import (
	"fmt"
	"math"
)

func rightJumpSearch(list []int, x int) int {
	b := int(math.Sqrt(float64(len(list))))                       // вычисляем размер блока(прыжка)
	start := 0                                                    // начальная позиция блока
	end := b - 1                                                  // конечная позиция блока
	for list[end] <= x && end < len(list)-1 && list[end+1] <= x { // пока конец блока меньше либо равно искомого элемента и в следующем блоке есть такой же элемент
		start = min(len(list)-1, start+b) // перемещаем начало блока вправо
		end = min(len(list)-1, end+b)     // перемещаем конец блока вправо
	}
	if x > list[end] { // если искомый элемент больше, чем последний элемент блока, значит не нашли нужный элемент
		return 0
	}
	for i := end; i >= start; i-- { // линейным поиском справа налево проходим по найденному блоку
		if list[i] == x { // если текущий элемент равен искомому, то возвращаем его индекс
			return i + 1
		}
	}
	return 0 // если дошли до сюда значит не нашли нужного элемента в списке
}

func makeArr(length int) []int {
	arr := make([]int, length)
	arrInt := make([]interface{}, length)

	for i := 0; i < length; i++ {
		arrInt[i] = &arr[i]
	}
	fmt.Scan(arrInt...)
	return arr
}

func main() {
	var N, M int

	fmt.Scan(&N, &M)

	arrN := makeArr(N)
	arrK := makeArr(M)

	for _, v := range arrK {
		fmt.Println(rightJumpSearch(arrN, v))
	}
}
