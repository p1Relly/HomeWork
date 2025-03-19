package main

import (
	"fmt"
)

type Student struct {
	Firstname    string
	Lastname     string
	Averagescore float64
}

func main() {
	var N int
	fmt.Scanln(&N)

	s := make([]Student, 0, N)

	for i := 0; i < N; i++ {
		fistname, lastname, informatics, mathematics, physics := "", "", 0.0, 0.0, 0.0
		fmt.Scan(&fistname, &lastname, &informatics, &mathematics, &physics)

		s = append(s, Student{
			Firstname:    fistname,
			Lastname:     lastname,
			Averagescore: (informatics + mathematics + physics) / 3,
		})

		//sort
		for i := len(s) - 1; i > 0; i-- {
			for j := 0; j < i; j++ {
				if s[j].Averagescore < s[j+1].Averagescore {
					s[j], s[j+1] = s[j+1], s[j]
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		fmt.Println(s[i].Firstname, s[i].Lastname)
	}
}
