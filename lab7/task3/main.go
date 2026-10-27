package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// Server хранит слушающий сокет и список активных соединений.
type Server struct {
	listener net.Listener
	wg       sync.WaitGroup
	mu       sync.Mutex
	conns    map[net.Conn]bool
}

func (s *Server) addConn(c net.Conn) {
	s.mu.Lock()
	s.conns[c] = true
	s.mu.Unlock()
}

func (s *Server) removeConn(c net.Conn) {
	s.mu.Lock()
	delete(s.conns, c)
	s.mu.Unlock()
}

// handle обслуживает одного клиента: читает строки и подтверждает каждую.
// Работает в отдельной горутине, поэтому клиенты не мешают друг другу.
func (s *Server) handle(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()
	defer s.removeConn(conn)

	addr := conn.RemoteAddr()
	fmt.Println("Клиент подключился:", addr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "exit" {
			break
		}
		fmt.Printf("[%s] %s\n", addr, text)
		fmt.Fprintln(conn, "Сообщение получено: "+text)
	}

	fmt.Println("Клиент отключился:", addr)
}

// closeAll предупреждает всех подключённых клиентов и закрывает их соединения.
func (s *Server) closeAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for c := range s.conns {
		fmt.Fprintln(c, "Сервер останавливается")
		c.Close()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Не удалось запустить сервер: ", err)
	}

	s := &Server{listener: listener, conns: make(map[net.Conn]bool)}

	// ctx отменяется, когда пользователь нажмёт Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// при сигнале закрываем слушающий сокет: Accept() вернёт ошибку и цикл закончится
	go func() {
		<-ctx.Done()
		fmt.Println("\nПолучен сигнал остановки, закрываю сервер...")
		listener.Close()
	}()

	fmt.Println("TCP-сервер слушает порт 9000 (остановка: Ctrl+C)...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			log.Println("Ошибка подключения:", err)
			continue
		}

		s.wg.Add(1)
		s.addConn(conn)
		go s.handle(conn)
	}

	// graceful shutdown: закрываем активные соединения и ждём завершения всех обработчиков
	s.closeAll()
	s.wg.Wait()
	fmt.Println("Сервер остановлен корректно")
}
