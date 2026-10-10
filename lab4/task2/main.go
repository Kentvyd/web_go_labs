package main

import "fmt"

// averageAge возвращает средний возраст всех людей в карте people.
func averageAge(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}

	sum := 0
	for _, age := range people {
		sum += age
	}
	return float64(sum) / float64(len(people))
}

func main() {
	ages := map[string]int{
		"Алексей": 25,
		"Мария":   30,
		"Иван":    22,
	}

	fmt.Println("Средний возраст:", averageAge(ages))
}
