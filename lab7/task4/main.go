package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// helloHandler отвечает на GET /hello приветствием.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Привет! Это HTTP-сервер на Go.")
}

// dataHandler принимает JSON в POST /data и выводит его содержимое в консоль.
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

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("HTTP-сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
