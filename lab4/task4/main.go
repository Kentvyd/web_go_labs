package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	fmt.Println("В верхнем регистре:", strings.ToUpper(line))
}
