package main

import (
	"errors"
	"fmt"
)

func main() {
	sorted_array := []int{1,2,3,4,8,15,22,57,99}
	targets := []int{1,4,22,99,100,9}

	for _, val := range targets {
		pos, err := binary_search(sorted_array, val)
		if (err != nil) {
			fmt.Printf("Value %d not found\n", val)
		} else {
			fmt.Printf("Value: %d, position: %d\n", val, pos)
		}
	}
}

func binary_search(array []int, value int) (int, error) {
	left := 0
	right := len(array) - 1

	for ;left <= right; {
		center := (left + right) / 2
		if (array[center] == value) {
			return center, nil
		} else if (array[center] < value) {
			left = center + 1
		} else if (array[center] > value) {
			right = center - 1
		}
	}

	return 0, errors.New("not found")
}
