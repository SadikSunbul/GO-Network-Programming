package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket bağlantısını yükseltmek için kullanılan upgrader nesnesi
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// İstemciden gelen veya gönderilen mesajı temsil eden yapı
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// Client, WebSocket bağlantısını ve mesaj göndermek için kullanılan kanalı temsil eder
type Client struct {
	conn *websocket.Conn
	send chan Message
}

// Hub, istemcileri yönetir ve mesajları yayınlar
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

// Yeni bir Hub oluşturur
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Hub'ı çalıştırır ve istemci yönetimini sağlar
func (h *Hub) run() {
	for {
		select {
		// Yeni bir istemci kaydedilir
		case client := <-h.register:
			h.clients[client] = true
		// Bir istemci kaydı silinir
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		// Bir mesaj yayınlanır
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// WebSocket bağlantısını yönetir ve istemciyi kaydeder
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{conn: conn, send: make(chan Message, 256)}
	hub.register <- client

	go client.writePump()
	go client.readPump(hub)
}

// İstemci için mesaj gönderme döngüsü
func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			json.NewEncoder(w).Encode(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// İstemci için mesaj okuma döngüsü
func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("error: %v", err)
			continue
		}
		hub.broadcast <- msg
	}
}

// Ana fonksiyon, sunucuyu başlatır ve WebSocket hizmetini sağlar
func main() {
	// JSON formatında mesajları işleyen bir WebSocket sunucusunu oluşturalım
	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
