package main

import (
	"fmt"

	"lab3/stringutils"
)

func main() {
	str := "Привет, Go!"

	fmt.Println("Исходная строка:", str)
	fmt.Println("Перевёрнутая строка:", stringutils.Reverse(str))
}
