package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// BankaHesabi, RPC yöntemlerini içeren bir yapıdır.
type BankaHesabi struct {
	Bakiye float64
}

// Sorgula, BankaHesabi yapısına ait bir yöntemdir ve mevcut bakiyeyi döndürür.
func (b *BankaHesabi) Sorgula(args *struct{}, yanit *float64) error {
	*yanit = b.Bakiye
	return nil
}

// Yatir, BankaHesabi yapısına ait bir yöntemdir ve belirtilen miktarda para yatırır.
func (b *BankaHesabi) Yatir(args *float64, yanit *struct{}) error {
	b.Bakiye += *args
	return nil
}

// Cek, BankaHesabi yapısına ait bir yöntemdir ve belirtilen miktarda para çeker.
func (b *BankaHesabi) Cek(args *float64, yanit *struct{}) error {
	if b.Bakiye < *args {
		return fmt.Errorf("Yetersiz bakiye")
	}
	b.Bakiye -= *args
	return nil
}

func main() {
	// BankaHesabi yapısını RPC sunucusuna kaydediyoruz.
	bankaHesabi := &BankaHesabi{Bakiye: 100.0}
	rpc.Register(bankaHesabi)

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
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

/*
err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
err = client.Call("BankaHesabi.Yatir", &yatir, &struct{}{})
err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
err = client.Call("BankaHesabi.Cek", &cek, &struct{}{})
err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
*/
