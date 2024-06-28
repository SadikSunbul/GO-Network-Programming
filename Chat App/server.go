package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var (
	clients   = make(map[string]net.Conn)
	broadcast = make(chan Message)
)

type Message struct {
	sender   string
	receiver string
	content  string
}

func main() {
	listener, err := net.Listen("tcp", ":12345") // Sunucuya bağlanıyoruz
	if err != nil {
		log.Fatalf("Sunucu başlatılamadı: %s", err)
	}
	defer listener.Close() // Sunucuyu kapatıyoruz

	log.Println("Sunucu 12345 portunda çalışıyor...")

	go handleMessages() // Mesajları dinliyoruz

	for {
		conn, err := listener.Accept() // Bir bağlantı alıyoruz
		if err != nil {
			log.Printf("Bağlantı kabul edilemedi: %s", err)
			continue
		}

		go handleConnection(conn) // Bağlantıyı işliyoruz
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Bağlantıyı kapatıyoruz

	reader := bufio.NewReader(conn) // Mesajları okuyoruz
	writer := bufio.NewWriter(conn) // Mesajları yazıyoruz

	writer.Flush()

	username, _ := reader.ReadString('\n') // Kullanıcı adını okuyoruz
	username = strings.TrimSpace(username) // Kullanıcı adını temizliyoruz

	clients[username] = conn // Kullanıcıyı map'e ekliyoruz

	log.Printf("%s bağlandı\n", username)

	for {
		message, err := reader.ReadString('\n') // Mesajı okuyoruz
		if err != nil {
			log.Printf("%s bağlantısı kesildi\n", username) // Bağlantıyı kapatıyoruz
			delete(clients, username)
			return
		}

		message = strings.TrimSpace(message)     // Mesajı temizliyoruz
		parts := strings.SplitN(message, " ", 2) // Mesajı boşluklara ayırıyoruz
		if len(parts) != 2 {                     // Mesajın boşluk sayısı 2 olmalı
			writer.WriteString("Geçersiz mesaj formatı\n")
			writer.Flush()
			continue
		}

		receiver := parts[0] // Mesajın alıcı bilgisini alıyoruz
		content := parts[1]  // Mesajın icerigini alıyoruz

		broadcast <- Message{sender: username, receiver: receiver, content: content} //
	}
}

// Mesajları dinliyoruz
func handleMessages() {
	for msg := range broadcast {
		if conn, ok := clients[msg.receiver]; ok {
			writer := bufio.NewWriter(conn)
			writer.WriteString(fmt.Sprintf("%s: %s\n", msg.sender, msg.content))
			writer.Flush()
		}
	}
}
