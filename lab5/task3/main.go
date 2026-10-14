package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// Area вычисляет площадь круга: π * r².
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5}

	fmt.Printf("Радиус: %.1f\n", c.radius)
	fmt.Printf("Площадь круга: %.2f\n", c.Area())
}
