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

func leftBinarySearch(arr []int, key int) int {
	left := -1
	right := len(arr) - 1
	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] < key {
			left = mid
		} else {
			right = mid
		}
	}
	if arr[right] == key {
		return right
	}
	return -1
}

func main() {
	var N, M int

	fmt.Scan(&N, &M)
	arrN := createArr(N)
	arrM := createArr(M)

	for i := range M {
		fmt.Println(leftBinarySearch(arrN, arrM[i]))
	}
}
