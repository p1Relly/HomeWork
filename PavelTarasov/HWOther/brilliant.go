package main

import (
	"fmt"
	"strings"
)

func printDiamond(n int) {
	symbol := []string{"#", "#"}
	mid := (n * 2) / 2
	var space, indent string

	fmt.Println("Мой бриллиант:")
	for i := 1; i < n*2; i++ {
		indent = strings.Repeat(" ", myAbs(mid-i))
		if i == 1 || i == n*2-1 {
			space = symbol[0]
		} else {
			space = strings.Join(symbol, strings.Repeat(" ", n*2-myAbs(mid-i)*2-3))
		}
		fmt.Println(indent + space)
	}
}

func myAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	printDiamond(3)
}
