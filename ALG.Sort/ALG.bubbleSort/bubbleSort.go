package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- { // на какое место будем ставим наибольший элемент
		flag := false // false - не было обменов, true - был хотя бы 1 обмен

		for j := 0; j < i; j++ { // проходим по не отсортированной последовательности
			if arr[j] > arr[j+1] { // если порядок элементов неправильный
				arr[j], arr[j+1] = arr[j+1], arr[j] // меняем местами пары
				flag = true                         // меняем флажок
			}
		}

		if flag == false { // если значение флага не поменялось
			return arr // массив отсортирован. Выходим с функции
		}
	}
	return arr
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
	var N int

	fmt.Scan(&N)
	arrN := makeArr(N)

	arrN = bubbleSort(arrN)
	for _, v := range arrN {
		fmt.Printf("%d ", v)
	}
}
