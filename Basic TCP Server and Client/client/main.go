package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080") // TCP sunucusuna bağlanıyor
	if err != nil {
		fmt.Println("Sunucuya bağlanılamadı:", err)
		return
	}
	defer conn.Close() // Bağlantıyı kapatıyor

	// Sunucuya mesaj gönderiliyor
	_, err = conn.Write([]byte("Merhaba, TCP!")) // TCP sunucusuna mesaj gönderiliyor
	if err != nil {
		fmt.Println("Mesaj gönderilemedi:", err)
		return
	}

	// Sunucudan yanıt okunuyor
	buffer := make([]byte, 1024) // 1024 byte'lık bir buffer oluştur
	n, err := conn.Read(buffer)  // Bağlantıdan gelen verileri buffer'a oku
	if err != nil {
		fmt.Println("Yanıt okunamadı:", err)
		return
	}
	// Okunan yanıt yazdırılıyor
	fmt.Println("Sunucudan gelen yanıt:", string(buffer[:n]))
}
