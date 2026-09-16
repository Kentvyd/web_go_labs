package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 5.5, height: 3.2}
	fmt.Println("Ширина:", rect.width)
	fmt.Println("Высота:", rect.height)
	fmt.Println("Площадь прямоугольника:", rect.Area())
}
