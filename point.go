package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (this *Point) Distance() float64 {
	return math.Pow(math.Pow(this.x, 2)+math.Pow(this.y, 2), 0.5)
}

func main() {
	a := Point{100, 100}
	fmt.Println("Distance = ", a.Distance())
}
