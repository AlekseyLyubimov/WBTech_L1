package main

import "fmt"

func main() {
	a, b := 69, 420
	fmt.Printf("Original order: %d, %d\n", a, b)
	a,b = b, a
	fmt.Printf("Flipped: %d, %d\n", a, b)
}