package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []interface{}{420, "string", true, make(chan int)}

	for _, d := range data {
		type_switch(d)
		type_reflect(d)
	}
}

func type_switch(i interface{}) {
	switch i.(type) {
		case int:
			println("int")
		case string:
			println("str")
		case bool:
			println("bool")
		case chan int:
			println("chan")
		default:
			println("unknown")
	}
}

func type_reflect(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
}