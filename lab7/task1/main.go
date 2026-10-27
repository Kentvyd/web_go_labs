package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}
	defer listener.Close()

	fmt.Println("TCP-сервер слушает порт 9000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка подключения:", err)
			continue
		}
		fmt.Println("Клиент подключился:", conn.RemoteAddr())
		handle(conn)
	}
}

// handle читает одно сообщение клиента, выводит его и отправляет подтверждение.
func handle(conn net.Conn) {
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	message = strings.TrimSpace(message)
	fmt.Println("Получено сообщение:", message)

	fmt.Fprintln(conn, "Сообщение получено: "+message)
}
