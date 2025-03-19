package main

import (
	"fmt"
	"math"
)

func leftJumpSearch(list []int, x int) int {
	b := int(math.Sqrt(float64(len(list)))) // вычисляем размер блока(прыжка)
	start := 0                              // начальная позиция блока
	end := b - 1                            // конечная позиция блока
	for list[end] < x {                     // пока конец блока меньше искомого элемента. Прервемся при равенстве
		if end == len(list)-1 { // если дошли до конца массива, выходим
			break
		}
		start = int(math.Min(float64(len(list)-1), float64(start+b))) // перемещаем начало блока вправо
		end = int(math.Min(float64(len(list)-1), float64(end+b)))     // перемещаем конец блокав право
	}
	if x > list[end] { // если искомый элемент больше, чем последний элемент блока, значит не нашли нужный элемент
		return 0
	}
	for i := start; i <= end; i++ { // линейным поиском слева направо проходим по найденному блоку
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
		fmt.Println(leftJumpSearch(arrN, v))
	}
}
