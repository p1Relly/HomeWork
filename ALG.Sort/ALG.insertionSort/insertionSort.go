package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	var buffer int
	for i := 1; i < len(arr); i++ {
		buffer = arr[i] // запоминаем в буфер элемент, который нужно вставить нужное место
		j := i          // индекс места, куда будем вставлять buffer
		// пока не дошли до начала массива и очередной элемент больше буфера
		for j > 0 && arr[j-1] > buffer {
			arr[j] = arr[j-1] // перемещаем больший элемент на одну позицию вправо
			j--               // передвигаем индекс для просмотра элемента, который стоит левее
		}
		arr[j] = buffer // место найдено, вставить элемент
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
	fmt.Scanln(&N)

	arrN := makeArr(N)

	arrN = insertionSort(arrN)

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", arrN[i])
	}
}
