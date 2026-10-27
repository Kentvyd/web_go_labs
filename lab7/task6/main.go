package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// страница чата, которую сервер отдаёт браузеру
const page = `<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="utf-8">
	<title>Чат</title>
</head>
<body>
	<h2>Чат на веб-сокетах</h2>
	<div id="chat" style="border:1px solid #999;height:300px;overflow:auto;padding:6px"></div>
	<input id="msg" placeholder="Сообщение" style="width:70%">
	<button onclick="sendMsg()">Отправить</button>
	<script>
		var name = prompt("Ваше имя:") || "Аноним";
		var ws = new WebSocket("ws://" + location.host + "/ws?name=" + encodeURIComponent(name));
		var chat = document.getElementById("chat");

		ws.onmessage = function (e) {
			var line = document.createElement("div");
			line.textContent = e.data;
			chat.appendChild(line);
			chat.scrollTop = chat.scrollHeight;
		};

		function sendMsg() {
			var input = document.getElementById("msg");
			if (input.value) {
				ws.send(input.value);
				input.value = "";
			}
		}

		document.getElementById("msg").addEventListener("keydown", function (e) {
			if (e.key === "Enter") {
				sendMsg();
			}
		});
	</script>
</body>
</html>`

// upgrader превращает обычное HTTP-соединение в веб-сокет.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Hub хранит подключённых клиентов и рассылает им сообщения.
type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func (h *Hub) add(c *websocket.Conn) {
	h.mu.Lock()
	h.clients[c] = true
	h.mu.Unlock()
}

func (h *Hub) remove(c *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, c)
	h.mu.Unlock()
}

// broadcast отправляет сообщение всем клиентам.
// Мьютекс гарантирует, что в каждое соединение пишет только одна горутина за раз.
func (h *Hub) broadcast(msg string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for c := range h.clients {
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			c.Close()
			delete(h.clients, c)
		}
	}
}

// wsHandler обслуживает одного клиента чата.
func (h *Hub) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка перехода на веб-сокет:", err)
		return
	}
	defer conn.Close()

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Аноним"
	}

	h.add(conn)
	h.broadcast("Участник " + name + " вошёл в чат")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		h.broadcast(name + ": " + string(msg))
	}

	h.remove(conn)
	h.broadcast("Участник " + name + " вышел из чата")
}

func main() {
	hub := &Hub{clients: make(map[*websocket.Conn]bool)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, page)
	})
	http.HandleFunc("/ws", hub.wsHandler)

	fmt.Println("Чат запущен: http://localhost:8080 (откройте страницу в двух вкладках)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
