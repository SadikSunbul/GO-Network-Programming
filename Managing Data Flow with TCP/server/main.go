package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
net.Listen fonksiyonu, Go programlama dilinde bir ağ soketi oluşturup dinlemeye başlamak için
kullanılır. Bu fonksiyon, belirtilen ağ adresi ve protokolü kullanarak bir dinleme soketi
(listener) oluşturur ve bu soketi kullanarak gelen bağlantıları kabul etmeye hazır hale getirir.
*/
func main() {
	listener, err := net.Listen("tcp", ":8080") // TCP sunucusu başlatılıyor
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
		return
	}

	defer listener.Close() // Sunucu kapatılıyor

	fmt.Println("Sunucu 8080 portunda dinlemede...")

	for {
		conn, err := listener.Accept() // Bir bağlantı kabul ediliyor
		if err != nil {
			fmt.Println("Bağlantı kabul edilemedi:", err)
			continue
		}

		go handleConnection(conn) // Her bağlantı için ayrı bir goroutine başlatılıyor
	}
}

// Bağlantıyı işleyen fonksiyon
func handleConnection(conn net.Conn) {
	defer conn.Close() // Bağlantıyı kapat

	// bufio ile veri akışını yönetme
	reader := bufio.NewReader(conn) // Veri akışını göstermek için bir reader oluştur
	for {
		// Satır satır okuma
		message, err := reader.ReadString('\n') // Satır satır okuma
		if err != nil {
			fmt.Println("Veri okunamadı:", err)
			return
		}

		fmt.Print("Alınan mesaj:", message) // Alınan mesajı ekrana yazdır

		// Yanıt gönderme
		_, err = conn.Write([]byte("Mesaj alındı\n")) // Mesajı sunucuya yolla
		if err != nil {
			fmt.Println("Yanıt gönderilemedi:", err)
			return
		}
	}
}
