package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP sunucusu başlatılıyor
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Sunucu başlatılamadı:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Sunucu 8080 portunda dinlemede...")

	for {
		// Bağlantıları kabul etme
		conn, err := listener.Accept() // Bir bağlantı kabul ediliyor
		if err != nil {
			fmt.Println("Bağlantı kabul edilemedi:", err)
			continue
		}

		// Her bağlantı için ayrı bir goroutine başlatılıyor
		go handleConnection(conn)
	}
}

// Bağlantıyı işleyen fonksiyon
func handleConnection(conn net.Conn) {
	defer conn.Close() // Bağlantıyı kapat

	for { //buradaki for sadece 1 den fazla kez mesaj gondermemizi saglar
		//veri oku
		buffer := make([]byte, 1024) // 1024 byte'lık bir buffer oluştur
		n, err := conn.Read(buffer)  // Bağlantıdan gelen verileri buffer'a oku
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Bağlantı kesildi")
				return
			}
			fmt.Println("Veri okunamadı:", err)
			return
		}
		//veriyi yazdır
		fmt.Println("Alınan mesaj:", string(buffer[:n]))

		// Bağlantıya yanıt gönderiliyor
		_, err = conn.Write([]byte("mesajınız kabul edildi"))
		if err != nil {
			fmt.Println("Yanıt gönderilemedi:", err)
			return
		}
	}
}
