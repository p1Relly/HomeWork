package main

import (
	"fmt"
	"math"
)

func jumpSearch(arr []int, key int) string {
	b := int(math.Sqrt(float64(len(arr))))
	start := 0
	end := b - 1
	for arr[end] < key {
		if end == len(arr)-1 {
			break
		}
		start = int(math.Min(float64(len(arr)-1), float64(start+b)))
		end = int(math.Min(float64(len(arr)-1), float64(end+b)))
	}
	if key > arr[end] {
		return "NO"
	}
	for i := end; i >= start; i-- {
		if arr[i] == key {
			return "YES"
		}
	}
	return "NO"
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

	for _, v := range arrK {
		fmt.Println(jumpSearch(arrN, v))
	}
}
