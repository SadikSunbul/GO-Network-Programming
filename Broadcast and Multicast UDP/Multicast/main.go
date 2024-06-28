package main

import (
	"fmt"
	"net"
)

/*
multicast, belirli bir grup cihaza veri göndermek için kullanılır.
*/
func main() {
	addr, err := net.ResolveUDPAddr("udp", "239.0.0.1:12345") // UDP adresi oluşturuluyor
	// Bu kod, UDP multicast mesajı gönderen bir istemci kodudur. Multicast mesajları, belirli bir IP adresi grubuna gönderilir. Kodda belirtilen 239.0.0.1 adresi, multicast grubu için kullanılan bir IP adresidir.
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

	_, err = conn.Write([]byte("Multicast message")) // UDP'ye mesaj yazılıyor
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}
	fmt.Println("Çok noktaya yayın mesajı gönderildi")
}
