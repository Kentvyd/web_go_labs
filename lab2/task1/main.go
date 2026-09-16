package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "- чётное")
	} else {
		fmt.Println(n, "- нечётное")
	}
}
