package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:12345") // Sunucuya bağlanıyoruz
	if err != nil {
		log.Fatalf("Sunucuya bağlanılamadı: %s", err) // Hata mesajını yazdırıyoruz
	}
	defer conn.Close() // Bağlantıyı kapatıyoruz

	reader := bufio.NewReader(conn) // Mesajları okuyoruz
	writer := bufio.NewWriter(conn) // Mesajları yazıyoruz

	go func() { // Mesajları ekrana yazdırıyoruz
		for { // Mesajları ekrana yazdırıyoruz
			message, err := reader.ReadString('\n') // Mesajı okuyoruz
			if err != nil {
				log.Fatalf("Mesaj alınamadı: %s", err)
			}
			fmt.Print(message)
		}
	}()
	fmt.Print("Kullanıcı adınızı girin: ")
	for { // Kullanıcıdan mesaj alıyoruz
		reader := bufio.NewReader(os.Stdin) // Kullanıcıdan mesaj alıyoruz
		fmt.Print("> ")
		text, _ := reader.ReadString('\n') // Mesajı okuyoruz
		text = strings.TrimSpace(text)     // Mesajı temizliyoruz

		writer.WriteString(text + "\n") // Mesajı yazıyoruz
		writer.Flush()                  // Mesajı yazdırıyoruz
	}
}
