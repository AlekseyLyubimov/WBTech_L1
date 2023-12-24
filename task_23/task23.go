package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(removeElement(nums, 2))
}


func removeElement(arr []int, pos int) ([]int, error) {
	if (pos >= 0 && pos <= len(arr)) {
		return append(arr[:pos], arr[pos+1:]...), nil
	} else {
		return nil, errors.New("out of bounds")
	}
}