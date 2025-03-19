package main

import (
	"fmt"
	"math"
)

func binarySearch(arr []int, key int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(arr) {
		left--
	} else if left < 0 {
		left = 0
	}
	return left
}

func searchNear(i, key int, arr []int) (indx int) {
	if i == len(arr)-1 {
		indx = checkDifference(key, arr, i-1, i)
	} else if i == 0 {
		indx = checkDifference(key, arr, i, i+1)
	} else {
		indx = checkDifference(key, arr, i-1, i, i+1)
	}
	return
}

func checkDifference(key int, arr []int, j ...int) (indx int) {
	minDifference := math.MaxInt

	for _, i := range j {
		if int(math.Abs(float64(key)-float64(arr[i]))) < minDifference {
			minDifference, indx = int(math.Abs(float64(key-arr[i]))), i
		}
	}

	return
}

func main() {
	var N, K int

	fmt.Scan(&N, &K)

	arrN := make([]int, N)
	arrNInt := make([]interface{}, N)
	arrK := make([]int, K)
	arrKInt := make([]interface{}, K)

	for i := 0; i < N; i++ {
		arrNInt[i] = &arrN[i]
	}
	fmt.Scan(arrNInt...)

	for i := 0; i < K; i++ {
		arrKInt[i] = &arrK[i]
	}
	fmt.Scan(arrKInt...)

	for _, key := range arrK {
		firstStep := binarySearch(arrN, key)

		secondStep := searchNear(firstStep, key, arrN)
		fmt.Println(arrN[secondStep])
	}
}
