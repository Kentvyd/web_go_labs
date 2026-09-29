package main

import "fmt"

func main() {
	words := []string{"go", "программирование", "код", "разработка", "функция"}

	longest := words[0]
	for _, w := range words {
		if len(w) > len(longest) {
			longest = w
		}
	}

	fmt.Println("Слова:", words)
	fmt.Println("Самое длинное слово:", longest)
}
