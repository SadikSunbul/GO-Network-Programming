package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	ctx := context.Background()

	// Google OAuth2 yapılandırması
	config := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		RedirectURL:  "YOUR_REDIRECT_URL",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// Yetkilendirme URL'si oluştur
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Yetkilendirme URL'si: %v\n", url)

	// Kullanıcıdan yetkilendirme kodunu al
	var code string
	fmt.Print("Yetkilendirme kodunu girin: ")
	fmt.Scan(&code)

	// Yetkilendirme kodunu kullanarak token al
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Token alınamadı: %v", err)
	}

	fmt.Printf("Token: %v\n", token)
}

/*
OAuth2, kullanıcıların üçüncü parti uygulamalara kimlik bilgilerini paylaşmadan belirli kaynaklara erişmesini sağlayan bir açık standarttır. OAuth2, web ve mobil uygulamaların kimlik doğrulama ve yetkilendirme işlemlerini basitleştirir.

### Neden OAuth2 Kullanılır?

1. **Güvenlik**: Kullanıcılar, hassas kimlik bilgilerini (kullanıcı adı ve parola gibi) doğrudan üçüncü parti uygulamalara paylaşmak yerine, sadece belirli izinleri yetkilendirirler. Bu, kullanıcı bilgilerinin daha güvenli bir şekilde korunmasını sağlar.
2. **Kullanım Kolaylığı**: Kullanıcılar, birden fazla hizmet için ayrı ayrı oturum açmak yerine, mevcut hesaplarını kullanarak hızlıca yetkilendirme yapabilirler.
3. **Esneklik**: OAuth2, farklı türde uygulamalar (web, mobil, masaüstü) ve farklı türde yetkilendirme akışları (yetkilendirme kodu, örtük, istemci kimlik bilgileri, cihaz kodu) için destek sunar.

### OAuth2 Nasıl Kullanılır?

OAuth2, genellikle aşağıdaki adımları içeren bir süreç izler:

1. **Uygulama Kaydı**: Üçüncü parti uygulama, OAuth2 sağlayıcısına (Google, Facebook, GitHub vb.) kaydolur ve bir istemci kimliği (client ID) ve istemci sırrı (client secret) alır.
2. **Yetkilendirme İsteği**: Uygulama, kullanıcıyı OAuth2 sağlayıcısının yetkilendirme sayfasına yönlendirir. Kullanıcı, uygulamanın hangi izinleri talep ettiğini görür ve onaylar.
3. **Yetkilendirme Kodu Alma**: Kullanıcı onayından sonra, OAuth2 sağlayıcısı, uygulamaya bir yetkilendirme kodu gönderir.
4. **Erişim Tokenı Alma**: Uygulama, yetkilendirme kodunu kullanarak OAuth2 sağlayıcısından bir erişim tokenı (access token) alır.
5. **API İstekleri**: Uygulama, erişim tokenını kullanarak korumalı kaynaklara (API'ler) erişir.

### Örnek Akış

1. **Uygulama Kaydı**:
   - Üçüncü parti uygulama, Google'a kaydolur ve bir istemci kimliği ve sırrı alır.

2. **Yetkilendirme İsteği**:
   - Kullanıcı, uygulamanın Google hesabına erişmesini isteyen bir istek görür ve onaylar.
   - Örnek URL: `https://accounts.google.com/o/oauth2/auth?client_id=CLIENT_ID&redirect_uri=REDIRECT_URI&response_type=code&scope=email%20profile`

3. **Yetkilendirme Kodu Alma**:
   - Kullanıcı onayından sonra, Google, uygulamaya bir yetkilendirme kodu gönderir.
   - Örnek URL: `https://example.com/callback?code=AUTHORIZATION_CODE`

4. **Erişim Tokenı Alma**:
   - Uygulama, yetkilendirme kodunu kullanarak Google'dan bir erişim tokenı alır.
   - Örnek İstek:
     ```http
     POST /oauth2/v4/token HTTP/1.1
     Host: accounts.google.com
     Content-Type: application/x-www-form-urlencoded

     code=AUTHORIZATION_CODE&client_id=CLIENT_ID&client_secret=CLIENT_SECRET&redirect_uri=REDIRECT_URI&grant_type=authorization_code
     ```

5. **API İstekleri**:
   - Uygulama, erişim tokenını kullanarak Google API'lerine istek gönderir.
   - Örnek İstek:
     ```http
     GET /plus/v1/people/me HTTP/1.1
     Host: www.googleapis.com
     Authorization: Bearer ACCESS_TOKEN
     ```

OAuth2, modern web ve mobil uygulamaların güvenli ve kullanıcı dostu bir şekilde kimlik doğrulama ve yetkilendirme işlemlerini gerçekleştirmesine olanak tanır.
*/
