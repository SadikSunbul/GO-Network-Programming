package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func main() {
	requestBody := RequestBody{Name: "John Doe", Age: 30} // Veri nesnesi oluşturuluyor
	jsonBody, err := json.Marshal(requestBody)            // JSON verisi oluşturuluyor
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/api", "application/json", bytes.NewBuffer(jsonBody)) // Web sayfası ile bağlantı kuruluyor
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Web sayfasını kapatıyor

	var responseBody ResponseBody                          // Web sayfasını okuyor
	err = json.NewDecoder(resp.Body).Decode(&responseBody) // Web sayfasını nesneye dönüştürüyor
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", responseBody.Message) // Web sayfasını ekrana yazdırıyor
}
