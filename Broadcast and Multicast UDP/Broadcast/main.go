package main

import (
	"fmt"
	"net"
)

/*
Broadcast, aynı ağdaki tüm cihazlara veri göndermek için kullanılır
test etmek için UDP serveri da çalıştırın
*/
func main() {
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:12345") // UDP adresi oluşturuluyor
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr) // UDP bağlantısı oluşturuluyor
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		return
	}
	defer conn.Close() // Bağlantıyı kapat

	_, err = conn.Write([]byte("Broadcast message")) // UDP'ye mesaj yazılıyor
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}
	fmt.Println("Yayın mesajı gönderildi")
}
