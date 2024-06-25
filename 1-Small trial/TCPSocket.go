package main

import (
	"fmt"
	"net"
)

func main() {
	//TCP soket oluşturma
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Localhost'ta sunucu dinleme:8080\n")

	for {
		//Bağlantıyı kabul etme
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Bağlantı kabul edilirken hata oluştu:", err.Error())
			return
		}
		fmt.Println("Baglanti kabul edildi")
		// Bağlantı üzerinden veri okuma ve yazma işlemleri yapılabilir
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()
	fmt.Println("Yeni bağlantı kuruldu")
	// Örnek: Bağlantıya "Hello, client" mesajı gönderme
	conn.Write([]byte("Hello, client\n"))

}
