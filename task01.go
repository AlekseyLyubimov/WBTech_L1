package main

import (
	"fmt"
)

func task01() {
	println("Task 1:")
	Action{Human{name: "Aleksey", age: 31}, "demo action"}.Greetings()
	pause()
}

type Human struct {
	age    int
	name   string
}

type Action struct {
	Human
	actionName string
}

func (human Human) Greetings() {
	fmt.Printf("Hi, my name is %s, I am %d years old.\n", human.name, human.age)
}