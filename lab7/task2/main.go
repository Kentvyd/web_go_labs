package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Print("Введите сообщение: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	// отправляем сообщение (одна строка, завершённая переводом строки)
	fmt.Fprintln(conn, text)

	// ждём ответ сервера
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}
	fmt.Println("Ответ сервера:", strings.TrimSpace(reply))

	// после выхода из main соединение закрывается (defer conn.Close())
}
