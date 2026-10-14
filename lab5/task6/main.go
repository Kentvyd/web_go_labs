package main

import "fmt"

// Stringer — всё, что умеет представить себя в виде строки.
type Stringer interface {
	String() string
}

type Book struct {
	title  string
	author string
	year   int
}

func (b Book) String() string {
	return fmt.Sprintf("«%s» — %s, %d г.", b.title, b.author, b.year)
}

func main() {
	b := Book{title: "Война и мир", author: "Лев Толстой", year: 1869}

	var s Stringer = b
	fmt.Println(s.String())

	// fmt.Println сама вызывает метод String(), если он есть
	fmt.Println(b)
}
