package main

import "fmt"

func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	fmt.Println("Среднее:", average(7, 4))
}
