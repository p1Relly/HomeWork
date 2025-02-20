package model

import (
	"fmt"
	"strings"
)

type Resident struct {
	Name    string
	Age     int
	Married bool
	Alive   bool
	Events  []string
}

func (r Resident) FlushInfo() string {
	return fmt.Sprintf("Житель %s (возраст: %d), статус: %s.\nСобытия: %s", r.Name, r.Age, printMarried(r.Married), printEvents(r.Events))
}

func (r Resident) Update() {

}

func printMarried(married bool) string {
	if married {
		return "в браке"
	}
	return "холост"
}

func printEvents(events []string) string {
	if len(events) == 0 {
		return "нет\n"
	}
	return fmt.Sprintf("\n%s\n\n", strings.Join(events, "\n"))
}
