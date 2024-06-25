package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	ls, err := net.Listen("tcp4", ":3000") // server listen port 3000 olarak ayarlandı
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ready!") // server bağlantısı hazır
	for {                            //1 DEN FAZLA client ustunden hızmet verebılmek ıcın yaptık
		conn, err := ls.Accept() // server bağlantısını kabul etti
		if err != nil {
			fmt.Print("connetion error:", err)
			continue
		}
		handler(conn) // client bağlantısı kabul edildi
	}
}

// client bağlantısı kabul edildi
func handler(conn net.Conn) {
	fmt.Println("connection accepted:", conn.RemoteAddr().String()) // client bağlantısı kabul edildi

	for { //1 DEN FAZLA client ustunden hızmet verebılmek ıcın yaptık
		header := make([]byte, 8)      // header hazır
		_, err := conn.Read(header[:]) // okuma yapıldı
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close() // client bağlantısı kapatıldı
			break
		}
		mlen := binary.LittleEndian.Uint32(header[4:]) // headerden uzunluk okundu
		databuff := make([]byte, mlen)                 // buffer hazır
		conn.Read(databuff[:])                         // okuma yapıldı
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close()
			break
		}

		var messagebug []byte
		messagebug = append(messagebug, header...)                     // header eklendi
		messagebug = append(messagebug, databuff...)                   // buffer eklendi
		mtype, _, msg := readMessage(messagebug)                       // mesaj okundu
		fmt.Printf("type:%d, len:%d, message:%s \n", mtype, mlen, msg) // mesaj yazdırıldı
	}
}

// go run main.go

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
