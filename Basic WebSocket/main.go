package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Web sayfası ile bağlantı kuruluyor
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close() // Web sayfasını kapatıyor

	for {
		messageType, message, err := conn.ReadMessage() // Web sayfasını okuyor
		if err != nil {
			log.Println("Read:", err)
			break
		}
		log.Printf("Received: %s", message) // Alınan mesajı ekrana yazdırıyor

		err = conn.WriteMessage(messageType, message) // Alınan mesajı geri gönderiyor
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)               // "/echo" yolunda WebSocket hizmeti sağlayacak işlevi tanımlıyor
	log.Println("Server started at :8080")       // Sunucunun 8080 numaralı bağlantı noktasında başlatıldığını belirtiyor
	log.Fatal(http.ListenAndServe(":8080", nil)) // Sunucuyu 8080 numaralı bağlantı noktasında başlatıyor
}

/*
Test için:

Postman'de WebSocket Bağlantısı Oluşturma:

Postman'i açın.

"New" butonuna tıklayın ve "WebSocket Request" seçeneğini seçin.

"Enter request URL" kısmına ws://localhost:8080/echo yazın.

"Connect" butonuna tıklayın.

Bağlantı kurulduktan sonra, mesaj gönderebilir ve sunucudan gelen yanıtı görebilirsiniz.

*/
