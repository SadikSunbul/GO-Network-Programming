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

	// RPC çağrısı için parametreleri hazırlıyoruz.
	args := []int{5, 10, 20}
	var yanit int

	// RPC çağrısı yapıyoruz.
	err = client.Call("Hesaplayici.Topla", &args, &yanit)
	if err != nil {
		fmt.Println("RPC hatası:", err)
		return
	}

	// Yanıtı yazdırıyoruz.
	fmt.Printf("Toplam: %d\n", yanit)
}

/*
HTTP (Hypertext Transfer Protocol):

Web uygulamalarında yaygın olarak kullanılan bir protokoldür.

İstemci-sunucu modeline dayalıdır ve istek-yanıt mekanizması sağlar.

HTTP, web sayfalarının yüklenmesi, API çağrıları ve diğer web tabanlı işlemler için kullanılır.

RPC (Remote Procedure Call):

Uzak bir makinedeki bir yordamı çağırmak için kullanılan bir tekniktir.

Programcının ağ detaylarını bilmeden uzak bir işlemi çağırmasını sağlar.

RPC, dağıtık sistemlerde ve ağ üzerinden işlemler gerçekleştirmek isteyen uygulamalarda kullanılır.

WebSocket:

İstemci ve sunucu arasında tam çift yönlü, etkileşimli iletişim sağlar.

Web uygulamalarında gerçek zamanlı iletişim için kullanılır (örneğin, sohbet uygulamaları, oyunlar vb.).

FTP (File Transfer Protocol):

Dosya transferi için kullanılan bir protokoldür.

İstemci ve sunucu arasında dosya yükleme ve indirme işlemlerini kolaylaştırır.

SMTP (Simple Mail Transfer Protocol):

E-posta göndermek için kullanılan bir protokoldür.

E-posta sunucuları arasında e-posta iletimi için kullanılır.

DNS (Domain Name System):

İnternet üzerindeki alan adlarını IP adreslerine çevirmek için kullanılan bir sistemdir.

Kullanıcıların alan adlarını kullanarak web sitelerine erişmesini sağlar.
*/
