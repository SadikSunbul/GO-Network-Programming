package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	// TCP üzerinden 3000 portuna bağlan
	conn, err := net.Dial("tcp4", ":3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Program sonunda bağlantıyı kapat

	// Bağlantı üzerinden gelen verileri dinlemek için yeni bir goroutine başlat
	go func() {
		for {
			buff := make([]byte, 8)      // 8 byte'lık bir buffer oluştur
			_, err := conn.Read(buff[:]) // Bağlantıdan gelen verileri buffer'a oku

			if err != nil {
				fmt.Println("Okuma hatası:", err)
				conn.Close() // Hata durumunda bağlantıyı kapat
				break        // Döngüden çık
			}
		}
	}()

	start := time.Now() // Zamanı başlat

	// 150000 kez dönecek bir döngü
	for i := 0; i < 10; i++ {
		// MessageTypeText türünde "hello from client" mesajı oluştur
		data := createMessage(MessageTypeText, "hello from client")
		// Mesajı bağlantı üzerinden gönder
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Yazma hatası:", err)
		}
	}

	end := time.Since(start) // Zamanı durdur ve geçen süreyi al
	fmt.Printf("Geçen süre: %s\n", end)

	// Sonsuz döngüde bekleyerek programın kapanmamasını sağla
	for {
		select {}
	}
}

const (
	MessageTypeJson = 1
	MessageTypeText = 2
	MessageTypeXML  = 3
)

/*
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string

	type	 | length  | data
*/
// mesaj oluşturuldu
func createMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))                        // Belirli bir boyutta bir buffer oluşturuldu
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))     // Mesajın türünü (mtype) buffer'a Little Endian formatında yazıldı (4 byte)
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data))) // Mesajın uzunluğunu (data uzunluğu) buffer'a Little Endian formatında yazıldı (4 byte)
	copy(buf[8:], []byte(data))                               // Mesajın veri kısmı (data) buffer'a kopyalandı (veri kısmı 8 byte'dan itibaren başlar)
	return buf                                                // Oluşturulan buffer fonksiyondan döndürüldü
}

// mesaj okundu
func readMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:]) // Buffer'ın başından itibaren 4 byte'lık kısım okunarak mesajın türü (mtype) elde edildi
	mlen = binary.LittleEndian.Uint32(data[4:])  // Buffer'ın 4. byte'ından itibaren 4 byte'lık kısım okunarak mesajın uzunluğu (mlen) elde edildi
	msg = string(data[8:])                       // Buffer'ın 8. byte'ından itibaren kalan kısmı string olarak alarak mesajın içeriği (msg) elde edildi
	return mtype, mlen, msg                      // Okunan değerler fonksiyondan döndürüldü
}
