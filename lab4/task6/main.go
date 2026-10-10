package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	numbers := make([]int, n)
	fmt.Println("Введите", n, "чисел:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	reversed := make([]int, n)
	for i, v := range numbers {
		reversed[n-1-i] = v
	}

	fmt.Println("Исходный массив:", numbers)
	fmt.Println("В обратном порядке:", reversed)
}
