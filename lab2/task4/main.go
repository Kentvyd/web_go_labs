package main

import (
	"fmt"
	"unicode/utf8"
)

func stringLength(s string) int {
	return len(s)
}

func main() {
	str := "Привет, Go!"

	fmt.Println("Строка:", str)
	fmt.Println("Длина в байтах (len):", stringLength(str))
	fmt.Println("Длина в символах (utf8.RuneCountInString):", utf8.RuneCountInString(str))
}
