package main

import "fmt"

//стандартный подход-отдельная функция на каждую операцию
func sum(x, y float64) float64 {
	return x + y
}

func diff(x, y float64) float64 {
	return x - y
}

func main() {
	// Вариант 2:
	s := sum(7.5, 2.3)
	d := diff(7.5, 2.3)

	fmt.Println("Сумма:", s, "Разность:", d)
}

// Go-функция с двумя возвращаемыми значениями
// func sumAndDiff(x, y float64) (float64, float64) {
// 	return x + y, x - y
// }

// func main() {
// 	// Вариант 1:
// 	s, d := sumAndDiff(7.5, 2.3)
//
// 	fmt.Println("Сумма:", s, "Разность:", d)
// }
