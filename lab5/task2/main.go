package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintInfo() {
	fmt.Println("Имя:", p.name, "| Возраст:", p.age)
}

// Birthday увеличивает возраст на 1 год.
// Получатель *Person (указатель) нужен, чтобы изменить самого человека, а не его копию.
func (p *Person) Birthday() {
	p.age++
}

func main() {
	p := Person{name: "Иван", age: 20}

	fmt.Println("До дня рождения:")
	p.PrintInfo()

	p.Birthday()

	fmt.Println("После дня рождения:")
	p.PrintInfo()
}
