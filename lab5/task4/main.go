package main

import (
	"fmt"
	"math"
)

// Shape — любая фигура, у которой можно посчитать площадь.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 4, height: 6}
	fmt.Printf("Площадь прямоугольника: %.2f\n", s.Area())

	s = Circle{radius: 5}
	fmt.Printf("Площадь круга: %.2f\n", s.Area())
}
