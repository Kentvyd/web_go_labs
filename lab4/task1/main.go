package main

import "fmt"

func main() {
	ages := map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    22,
	}

	// добавление нового человека
	ages["Ольга"] = 28

	fmt.Println("Все записи:")
	for name, age := range ages {
		fmt.Println(name, "-", age)
	}
}
