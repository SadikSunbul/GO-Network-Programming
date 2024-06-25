package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp4", ":3000") // server listen port 3000 olarak ayarlandı
	if err != nil {
		panic(err)
	}

	defer conn.Close() // server bağlantısı kapatıldı

	go func() { //1 DEN FAZLA client ustunden hızmet verebılmek ıcın yaptık
		for {
			buff := make([]byte, 8)        // buffer hazır
			_, err := conn.Read(buff[:])   // okuma yapıldı
			_, _, msg := readMessage(buff) // mesaj okundu
			if err != nil {
				fmt.Println("read error:", err)
				conn.Close() // client bağlantısı kapatıldı
				break
			}
			fmt.Printf("mssage from server : %s \n", msg) // mesaj yazdırıldı
		}
	}()

	for i := 0; i < 10; i++ {
		data := createMessage(MessageTypeText, "hello from client") // mesaj olusturuldu
		_, err = conn.Write(data)                                   // mesaj yazıldı
		if err != nil {
			fmt.Println("write error:", err)
		}
	}

	for {
		select {}
	}

}

const (
	MessageTypeJson = 1
	MessageTypeText = 2
	MessageTypeXML  = 3
)

/*
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string

	type	 | length  | data
*/
// mesaj olusturuldu
func createMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))                        // buffer olusturuldu
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))     // 0 dan ıtıbaren
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data))) //4 en ıtıbaren
	copy(buf[8:], []byte(data))                               //8 den ıtıbaren
	return buf
}

// mesaj okundu
func readMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:]) // 0 dan ıtıbaren
	mlen = binary.LittleEndian.Uint32(data[4:])  //4 en ıtıbaren
	msg = string(data[8:])                       // 8 den ıtıbaren
	return mtype, mlen, msg
}
