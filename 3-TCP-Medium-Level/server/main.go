package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime/pprof"
)

func main() {
	// cpu.out adında bir dosya oluşturuluyor
	f, err := os.Create("cpu.out")
	if err != nil {
		panic(err) // Eğer dosya oluşturulamazsa programı sonlandır
	}
	defer f.Close() // Dosya işlemi tamamlandıktan sonra dosyayı kapat

	// CPU profilini oluşturulmuş dosyaya yazmaya başla
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile() // CPU profilini dosyaya yazmayı durdur

	// TCP4 üzerinde 3000 portunu dinlemeye başla
	ls, err := net.Listen("tcp4", ":3000")
	if err != nil {
		panic(err) // Eğer bağlantı noktası dinlenemiyorsa programı sonlandır
	}

	fmt.Println("Bağlantı hazır!") // Bağlantı noktası dinlenmeye hazır mesajı yazdır

	// Bağlantıları kabul etmek için yeni bir goroutine başlat
	go func() {
		for {
			// Yeni bir istemci bağlantısı kabul edilene kadar döngüde kal
			conn, err := ls.Accept()
			if err != nil {
				fmt.Println("Bağlantı hatası:", err)
				continue // Hata durumunda döngüyü tekrar başlat
			}
			// Bağlantıyı işleyecek olan fonksiyonu çağır
			handler(conn)
		}
	}()

	// Program sonlandırılana kadar kesilme (interrupt) sinyali bekleniyor
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Kesilme sinyali alınıncaya kadar bekleniyor
}

// client bağlantısı kabul edildi
func handler(conn net.Conn) {
	fmt.Println("Bağlantı kabul edildi:", conn.RemoteAddr().String())

	for {

		header := make([]byte, 8)      // 8 byte'lık bir header dizisi oluşturuldu
		_, err := conn.Read(header[:]) // Bağlantıdan gelen veriler header dizisine okundu
		if err != nil {
			fmt.Println("Okuma hatası:", err)
			conn.Close() // Hata durumunda bağlantı kapatıldı
			break        // Döngüden çıkıldı
		}

		// Header'dan mesajın uzunluğu (mlen) okundu
		mlen := binary.LittleEndian.Uint32(header[4:])

		// Mesajın tamamını almak için mlen uzunluğunda bir data buffer oluşturuldu
		databuff := make([]byte, mlen)
		_, err = conn.Read(databuff[:]) // Bağlantıdan gelen veriler data buffer'a okundu
		if err != nil {
			fmt.Println("Okuma hatası:", err)
			conn.Close() // Hata durumunda bağlantı kapatıldı
			break        // Döngüden çıkıldı
		}

		// Header ile data buffer birleştirilerek messagebug dizisine aktarıldı
		var messagebug []byte
		messagebug = append(messagebug, header...)
		messagebug = append(messagebug, databuff...)

		// readMessage fonksiyonu ile messagebug dizisi işlenerek mesajın türü (mtype), uzunluğu (mlen) ve içeriği (msg) elde edildi
		mtype, mlen, msg := readMessage(messagebug)
		fmt.Printf("Tür:%d, Uzunluk:%d, Mesaj:%s \n", mtype, mlen, msg)
	}
}

// go run main.go

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
