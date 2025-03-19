package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	str := strings.Fields(scanner.Text())

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, N)
	for i := range arr {
		num, _ := strconv.Atoi(str[i])

		arr[i] = num
	}

	maxX := math.MaxInt
	newX, _ := strconv.Atoi(str[0])
	for i := range N {
		if int(math.Abs(float64(x)-math.Abs(float64(arr[i])))) <= maxX {
			maxX, newX = int(math.Abs(float64(x-arr[i]))), arr[i]
		}
	}
	fmt.Println(newX)
}
