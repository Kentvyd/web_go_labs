package main

import (
	"fmt"
	"math"
)

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

// printAreas выводит тип и площадь каждой фигуры из среза.
func printAreas(shapes []Shape) {
	for _, s := range shapes {
		fmt.Printf("%T: %.2f\n", s, s.Area())
	}
}

func main() {
	shapes := []Shape{
		Rectangle{width: 4, height: 6},
		Circle{radius: 5},
		Rectangle{width: 2, height: 3},
	}

	printAreas(shapes)
}
