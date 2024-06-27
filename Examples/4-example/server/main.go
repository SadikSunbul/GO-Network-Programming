package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	Listen()

}

func Connection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Yeni bağlantı kuruldu")

	a := deneme{Isim: "sadık", Soyisim: "sünbül"}
	d, _ := json.Marshal(a)

	conn.Write(d)
}

type deneme struct {
	Isim    string `json:"isim"`
	Soyisim string `json:"soyisim"`
}

func Listen() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Localhost'ta sunucu dinleme:8080\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Bağlantı kabul edilirken hata oluştu:", err.Error())
			return
		}
		fmt.Println("Baglanti kabul edildi")
		go Connection(conn)
	}
}
