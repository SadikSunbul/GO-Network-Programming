package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Web sayfasını ekrana yazdır
	
}

func main() {
	http.HandleFunc("/", helloHandler) // Web sayfası oluştur
	fmt.Println("Sunucuyu 8080 numaralı bağlantı noktasından başlatma\n")
	if err := http.ListenAndServe(":8080", nil); err != nil { // Sunucuyu 8080 numaralı bağlantı noktasından başlat
		fmt.Println(err)
	}
}
