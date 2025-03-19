package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func arrConvert(str []string, Len int) []int {
	arr := make([]int, Len)

	for i := range str {
		num, err := strconv.Atoi(str[i])
		if err != nil {
			log.Fatal(err)
		}

		arr[i] = num
	}
	return arr
}

func binarySearch(arr []int, num int) string {
	left := 0
	right := len(arr) - 1
	for left <= right {
		m := left + (right-left)/2 // вычисление серединного элемента
		if arr[m] == num {         // если серединный элемент равен ключу, то выводим индекс серединного элемента
			return "YES"
		}
		if arr[m] < num { // если серединный элемент меньше ключа, то смещаем левую границу
			left = m + 1
		} else {
			right = m - 1 // если серединный элемент больше ключа, то смещаем правую границу
		}
	}
	return "NO" // если не нашли соответствующего элемента, возвращаем "-1"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	NK := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(NK[0])
	K, _ := strconv.Atoi(NK[1])

	scanner.Scan()
	strN := strings.Fields(scanner.Text())

	scanner.Scan()
	strK := strings.Fields(scanner.Text())

	arrN := arrConvert(strN, N)
	arrK := arrConvert(strK, K)

	for _, num := range arrK {
		fmt.Println(binarySearch(arrN, num))
	}
}
