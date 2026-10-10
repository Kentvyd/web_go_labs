package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите числа через пробел: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Fields(line)

	sum := 0
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		sum += n
	}

	fmt.Println("Сумма:", sum)
}
