package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080") // Web sayfası ile bağlantı kuruluyor
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Web sayfasını kapatıyor

	body, err := ioutil.ReadAll(resp.Body) // Web sayfasını okuyor
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(body)) // Web sayfasını ekrana yazdırıyor
}
