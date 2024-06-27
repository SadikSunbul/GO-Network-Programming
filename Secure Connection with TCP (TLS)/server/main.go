package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key") // Sertifika yükleniyor
	if err != nil {
		log.Fatalf("Sunucu sertifikası yüklenemedi: %s", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}} // Sertifika eklendiyor
	listener, err := tls.Listen("tcp", ":8443", config)          // TCP sunucusu başlatılıyor
	if err != nil {
		log.Fatalf("TLS sunucusu başlatılamadı: %s", err)
	}

	defer listener.Close() // Sunucu kapatılıyor

	fmt.Println("TLS sunucusu 8443 portunda dinlemede...")

	for {
		conn, err := listener.Accept() // Bir bağlantı kabul ediliyor
		if err != nil {
			fmt.Println("Bağlantı kabul edilemedi:", err)
			continue
		}

		go handleConnection(conn) // Her bağlantı için ayrı bir goroutine başlatılıyor
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Bağlantıyı kapat

	buffer := make([]byte, 1024) // 1024 byte'lık bir buffer oluştur
	n, err := conn.Read(buffer)  // Bağlantıdan gelen verileri buffer'a oku

	if err != nil {
		fmt.Println("Veri okunamadı:", err)
		return
	}

	fmt.Println("Alınan mesaj:", string(buffer[:n])) // Alınan mesajı ekrana yazdır

	_, err = conn.Write([]byte("Mesaj alındı")) // Mesajı sunucuya yolla
	if err != nil {
		fmt.Println("Yanıt gönderilemedi:", err)
		return
	}
}
