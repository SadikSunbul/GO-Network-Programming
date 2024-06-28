package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT için özel claim
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// JWT oluştur
	tokenString, err := createJWT()
	if err != nil {
		log.Fatalf("JWT oluşturulamadı: %v", err)
	}
	fmt.Printf("JWT: %v\n", tokenString)

	// JWT doğrula
	claims, err := validateJWT(tokenString)
	if err != nil {
		log.Fatalf("JWT doğrulanamadı: %v", err)
	}
	fmt.Printf("Claims: %v\n", claims)
}

func createJWT() (string, error) {
	// Özel claim oluştur
	claims := CustomClaims{
		Username: "exampleUser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "userAuthentication",
		},
	}

	// JWT oluştur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("your-256-bit-secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-256-bit-secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("geçersiz token")
}
