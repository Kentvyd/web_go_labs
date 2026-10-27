package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Привет! Это HTTP-сервер на Go.")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Получены данные:", data)
	fmt.Fprintln(w, "Данные получены")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Текущее время:", time.Now().Format("2006-01-02 15:04:05"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Лабораторная работа 7: маршрутизация и middleware")
}

// loggingMiddleware оборачивает обработчик и после его работы
// записывает в лог метод, URL и время выполнения запроса.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("HTTP-сервер с middleware запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
}
