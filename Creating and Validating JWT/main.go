package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT için özel claim yapısı
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// JWT oluşturma fonksiyonunu çağır ve token al
	tokenString, err := createJWT()
	if err != nil {
		log.Fatalf("JWT oluşturulamadı: %v", err)
	}
	fmt.Printf("JWT: %v\n", tokenString)

	// JWT doğrulama fonksiyonunu çağır ve claim bilgilerini al
	claims, err := validateJWT(tokenString)
	if err != nil {
		log.Fatalf("JWT doğrulanamadı: %v", err)
	}
	fmt.Printf("Claims: %v\n", claims)
}

// JWT oluşturma fonksiyonu
func createJWT() (string, error) {
	// Özel claim bilgilerini oluştur
	claims := CustomClaims{
		Username: "exampleUser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token'ın süresi 24 saat sonra dolar
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Token'ın oluşturulma zamanı
			Subject:   "userAuthentication",                               // Token'ın konusu
		},
	}

	// JWT oluştur ve HS256 ile şifrele
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-256-bit-secret")) // Token'ı gizli anahtarla imzala
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWT doğrulama fonksiyonu
func validateJWT(tokenString string) (*CustomClaims, error) {
	// Token'ı parse et ve claim bilgilerini al
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-256-bit-secret"), nil // Token'ı doğrulamak için gizli anahtarı kullan
	})
	if err != nil {
		return nil, err
	}

	// Token geçerli ise claim bilgilerini döndür
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("geçersiz token")
}
