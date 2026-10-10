package main

import "fmt"

func main() {
	ages := map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    22,
	}

	fmt.Println("До удаления:", ages)

	delete(ages, "Мария")

	fmt.Println("После удаления записи \"Мария\":", ages)
}
