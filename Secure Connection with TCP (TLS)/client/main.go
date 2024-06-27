package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	config := &tls.Config{ // Sertifika eklendi
		InsecureSkipVerify: true, // Sertifika doğrulamasını atla (test için)
	}

	conn, err := tls.Dial("tcp", "localhost:8443", config) // TCP sunucusuna bağlanıyor
	if err != nil {
		log.Fatalf("Sunucuya bağlanılamadı: %s", err) // Hata mesajı ekrana yazdırılıyor
	}
	defer conn.Close() // Bağlantıyı kapatıyor

	_, err = conn.Write([]byte("Merhaba, TLS!")) // Mesajı sunucuya yolla
	if err != nil {
		log.Fatalf("Mesaj gönderilemedi: %s", err) // Hata mesajı ekrana yazdırılıyor
	}

	buffer := make([]byte, 1024) // 1024 byte'lık bir buffer oluştur
	n, err := conn.Read(buffer)  // Bağlantıdan gelen verileri buffer'a oku
	if err != nil {
		log.Fatalf("Yanıt okunamadı: %s", err) // Hata mesajı ekrana yazdırılıyor
	}

	fmt.Println("Sunucudan gelen yanıt:", string(buffer[:n])) // Alınan mesajı ekrana yazdır

}
