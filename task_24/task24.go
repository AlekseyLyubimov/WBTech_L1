package main

import (
	"math"
)

func main() {

	pt1 := Point{0, 0}
	pt2 := Point{100, 0}
	println(pt1.DistanceTo(&pt2))

	pt3 := Point{50, 250}
	pt4 := Point{50, -250}
	println(pt3.DistanceTo(&pt4))
}

type Point struct {
	x float64
	y float64
}

func (start *Point) DistanceTo(dest *Point) float64 {
	x, y := start.x-dest.x, start.y-dest.y
	return math.Sqrt(x*x + y*y) 
}
