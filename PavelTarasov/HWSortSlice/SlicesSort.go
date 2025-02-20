package main

import (
	"fmt"
	"slices"
)

func swapArr(arrFirst []int, arrSecond []int) ([]int, []int) {
	arrTemporary := make([]int, len(arrFirst))
	arrTemporary = arrFirst[:len(arrFirst)]

	arrFirst = arrSecond[:len(arrSecond)]
	arrSecond = arrTemporary[:len(arrTemporary)]

	slices.Clip(arrFirst)
	slices.Clip(arrSecond)
	return arrFirst, arrSecond
}

func sumArr(arr []int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func magicSort(arr [][]int) {
	for range len(arr) - 1 {
		for i := range len(arr) - 1 {
			if sumArr(arr[i]) > sumArr(arr[i+1]) {
				arr[i], arr[i+1] = swapArr(arr[i], arr[i+1])
			}
		}
	}

	for _, row := range arr {
		slices.SortFunc(row, func(a, b int) int {
			if a == 0 && b != 0 {
				return -1
			}
			if a != 0 && b == 0 {
				return 1
			}
			if a%2 == 0 && b%2 != 0 {
				return -1
			}
			if a%2 != 0 && b%2 == 0 {
				return 1
			}
			if a > b {
				return -1
			}
			if b > a {
				return 1
			}
			return 0
		})
	}
}

func main() {
	arr := [][]int{{3, 1, 4, 1}, {2, 2, 2}, {5, 0, 6, 3, -8, 1}, {4, 6, 8, 2}}
	magicSort(arr)
	fmt.Println(arr)
}
