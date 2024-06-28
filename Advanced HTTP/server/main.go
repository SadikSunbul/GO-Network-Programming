package main

import (
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

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // POST isteği kontrolü
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody) // JSON verisi okunuyor
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest) // JSON verisi geçersizse hata veriliyor
		return
	}

	responseBody := ResponseBody{
		Message: fmt.Sprintf("Hello, %s! You are %d years old.", requestBody.Name, requestBody.Age), // Mesaj oluşturuluyor
	}

	w.Header().Set("Content-Type", "application/json") // veri tipi ayarlanıyor
	json.NewEncoder(w).Encode(responseBody)            // veri gönderiliyor
}

func main() {
	http.HandleFunc("/api", apiHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
