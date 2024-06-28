package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Hesaplayici, RPC yöntemlerini içeren bir yapıdır.
type Hesaplayici struct{}

// Topla, Hesaplayici yapısına ait bir yöntemdir ve iki sayıyı toplar.
func (h *Hesaplayici) Topla(args *[]int, yanit *int) error {
	*yanit = 0
	for _, sayi := range *args {
		*yanit += sayi
	}
	return nil
}

func main() {
	// Hesaplayici yapısını RPC sunucusuna kaydediyoruz.
	hesaplayici := new(Hesaplayici)
	rpc.Register(hesaplayici)

	// TCP bağlantı noktasını dinlemeye başlıyoruz.
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Dinleme hatası:", err)
	}
	defer listener.Close()

	log.Println("Sunucu 1234 portunda dinlemede")

	// Bağlantıları kabul ediyoruz ve her bağlantı için yeni bir goroutine başlatıyoruz.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Kabul hatası:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // JSON RPC istemcisi oluşturuyoruz.
	}
}
