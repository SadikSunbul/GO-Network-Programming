package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*
net.Dial fonksiyonu, Go programlama dilinde ağ bağlantıları oluşturmak için kullanılır.
Bu fonksiyon, belirtilen ağ adresine ve protokole göre bir bağlantı kurmayı amaçlar.
*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // TCP sunucusuna bağlanıyor
	if err != nil {
		fmt.Println("Sunucuya bağlanılamadı:", err)
		return
	}
	defer conn.Close() // Bağlantıyı kapatıyor

	reader := bufio.NewReader(os.Stdin) // Konsol'dan mesaj almak için bir reader oluştur

	for {
		fmt.Print("Mesajınız: ")
		message, _ := reader.ReadString('\n')

		// Sunucuya mesaj gönderme
		_, err = conn.Write([]byte(message)) // Mesajı sunucuya yolla
		if err != nil {
			fmt.Println("Mesaj gönderilemedi:", err)
			return
		}

		// Sunucudan yanıt okuma
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer) // Bağlantıdan gelen verileri buffer'a oku
		if err != nil {
			fmt.Println("Yanıt okunamadı:", err)
			return
		}
		fmt.Print("Sunucudan gelen yanıt:", string(buffer[:n]))
	}
}
