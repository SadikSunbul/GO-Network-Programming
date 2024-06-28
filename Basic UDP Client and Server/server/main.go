package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":12345") // UDP adresi oluşturuluyor
	if err != nil {
		fmt.Println("UDP adresi çözümlenirken hata oluştu:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr) // UDP bağlantısı oluşturuluyor
	if err != nil {
		fmt.Println("UDP'de dinleme hatası:", err)
		return
	}
	defer conn.Close() // Bağlantıyı kapat

	fmt.Println("UDP sunucusu dinleme tarihi:12345")

	buffer := make([]byte, 1024) // 1024 byte'lık bir buffer oluştur
	for {
		n, addr, err := conn.ReadFromUDP(buffer) // UDP'den gelen verileri buffer'a oku
		if err != nil {
			fmt.Println("UDP'den okuma hatası:", err)
			continue
		}
		fmt.Printf("Kabul edilmiş %d gelen bayt %s: %s\n", n, addr, string(buffer[:n])) // Alınan baytları ekrana yazdır

		_, err = conn.WriteToUDP([]byte("Message received"), addr) // Alınan baytları sunucuya yolla
		if err != nil {
			fmt.Println("UDP'ye yazarken hata oluştu:", err)
		}
	}

}
