package main

import (
	"fmt"

	"lab3/mathutils"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Println("Факториал числа", n, "=", mathutils.Factorial(n))
}
