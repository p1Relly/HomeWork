package main

import (
	"fmt"
)

func createArr(length int) []int {
	arr := make([]int, length)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func rightBinarySearch(arr []int, key int) int {
	left := 0
	right := len(arr)
	for left+1 < right {
		m := left + (right-left)/2
		if arr[m] <= key {
			left = m
		} else {
			right = m
		}
	}
	if arr[left] == key {
		return left
	}
	return -1
}

func main() {
	var N, M int

	fmt.Scan(&N, &M)
	arrN := createArr(N)
	arrM := createArr(M)

	for i := range M {
		fmt.Println(rightBinarySearch(arrN, arrM[i]))
	}
}
