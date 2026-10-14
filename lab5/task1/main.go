package main

import "fmt"

type Person struct {
	name string
	age  int
}

// PrintInfo выводит информацию о человеке.
func (p Person) PrintInfo() {
	fmt.Println("Имя:", p.name, "| Возраст:", p.age)
}

func main() {
	p := Person{name: "Иван", age: 20}
	p.PrintInfo()
}
