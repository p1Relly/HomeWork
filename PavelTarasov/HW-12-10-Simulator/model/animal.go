package model

import (
	"fmt"
)

type Animal struct {
	Name   string
	Age    int
	Type   string
	Alive  bool
	Events []string
}

func (a Animal) FlushInfo() string {
	return fmt.Sprintf("Животное %s (%s, возраст: %d).\nСобытия: %s", a.Name, a.Type, a.Age, printEvents(a.Events))
}

func (a Animal) Update() {
}
