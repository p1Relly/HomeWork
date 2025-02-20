package main

import (
	"fmt"
)

func main() {
	var j, s string
	mJ := map[rune]struct{}{}

	fmt.Scan(&j, &s)

	for _, char := range j {
		mJ[char] = struct{}{}
	}

	result := 0
	for _, char := range s {
		_, ok := mJ[char]
		if ok {
			result++
		}
	}

	fmt.Println(result)
}

/*

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str := strings.Split(scanner.Text(), " ")

	sum := 0

	for _, value := range str {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		sum += int(num)
	}
	fmt.Println(sum)
}

*/
