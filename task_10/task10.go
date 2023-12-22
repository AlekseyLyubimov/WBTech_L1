package main

import (
	"fmt"
	"math"
)

func main() {
	vals := []float32{-25.4, -15.4, -9.4, 9.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	val_map := make(map[int][]float32)

	for _, val := range vals {
		pos := int(math.Floor(float64(val/10))) * 10
		val_map[pos] = append(val_map[pos], val)
	}

	fmt.Println(val_map)
}
