package main

//counter

import (
	"fmt"
	"slices"
)

func counterSort(arr []int, N int) {
	k := slices.Max(arr) // ширина используемого диапазона
	counts := make([]int, k+1)

	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	for i := k; i >= 0; i-- {
		for j := 0; j < counts[i]; j++ { // counts[i]-количество чисел i
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var N int
	fmt.Scanln(&N)

	arrN := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arrN[i])
	}

	counterSort(arrN, N)
}
