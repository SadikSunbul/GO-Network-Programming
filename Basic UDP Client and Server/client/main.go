package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:12345") // UDP adresi oluşturuluyor
	if err != nil {
		fmt.Println("UDP adresi çözümlenirken hata oluştu:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr) // UDP bağlantısı oluşturuluyor
	if err != nil {
		fmt.Println("UDP'yi çevirirken hata oluştu:", err)
		return
	}
	defer conn.Close() // Bağlantıyı kapat

	_, err = conn.Write([]byte("Merhaba UDP sunucusu!")) // UDP'ye mesaj yazılıyor
	if err != nil {
		fmt.Println("UDP'ye yazarken hata oluştu:", err)
		return
	}

	buffer := make([]byte, 1024)          // 1024 byte'lık bir buffer oluştur
	n, _, err := conn.ReadFromUDP(buffer) // UDP'den gelen verileri buffer'a oku
	if err != nil {
		fmt.Println("UDP'den okuma hatası:", err)
		return
	}
	fmt.Printf("Sunucudan alındı: %s\n", string(buffer[:n])) // Alınan baytları ekrana yazdır
}
