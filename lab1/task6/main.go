package main

import "fmt"

func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	fmt.Println("cреднее:", average(4, 8, 15))
}
