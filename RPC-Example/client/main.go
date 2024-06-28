package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	// Sunucuya TCP bağlantısı açıyoruz.
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Bağlantı hatası:", err)
		return
	}
	defer conn.Close()

	// JSON RPC istemcisi oluşturuyoruz.
	client := jsonrpc.NewClient(conn)

	// Bakiye sorgulama
	var bakiye float64
	err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
	if err != nil {
		fmt.Println("Bakiye sorgulama hatası:", err)
		return
	}
	fmt.Printf("Mevcut bakiye: %.2f\n", bakiye)

	// Para yatırma
	var yatir float64 = 50.0
	err = client.Call("BankaHesabi.Yatir", &yatir, &struct{}{})
	if err != nil {
		fmt.Println("Para yatırma hatası:", err)
		return
	}
	fmt.Printf("%.2f para yatırıldı\n", yatir)

	// Bakiye sorgulama
	err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
	if err != nil {
		fmt.Println("Bakiye sorgulama hatası:", err)
		return
	}
	fmt.Printf("Mevcut bakiye: %.2f\n", bakiye)

	// Para çekme
	var cek float64 = 75.0
	err = client.Call("BankaHesabi.Cek", &cek, &struct{}{})
	if err != nil {
		fmt.Println("Para çekme hatası:", err)
		return
	}
	fmt.Printf("%.2f para çekildi\n", cek)

	// Bakiye sorgulama
	err = client.Call("BankaHesabi.Sorgula", &struct{}{}, &bakiye)
	if err != nil {
		fmt.Println("Bakiye sorgulama hatası:", err)
		return
	}
	fmt.Printf("Mevcut bakiye: %.2f\n", bakiye)
}
