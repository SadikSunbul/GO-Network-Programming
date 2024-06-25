package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	ls, err := net.Listen("tcp4", ":3000")
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ready!")
	for { //1 DEN FAZLA client ustunden hızmet verebılmek ıcın yaptık
		conn, err := ls.Accept()
		if err != nil {
			fmt.Print("connetion error:", err)
			continue
		}
		handler(conn)
	}
}

func handler(conn net.Conn) {
	fmt.Println("connection accepted:", conn.RemoteAddr().String())

	for {
		header := make([]byte, 8)
		_, err := conn.Read(header[:])
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close()
			break
		}
		mlen := binary.LittleEndian.Uint32(header[4:])
		databuff := make([]byte, mlen)
		conn.Read(databuff[:])
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close()
			break
		}

		var messagebug []byte
		messagebug = append(messagebug, header...)
		messagebug = append(messagebug, databuff...)
		mtype, _, msg := readMessage(messagebug)
		fmt.Printf("type:%d, len:%d, message:%s \n", mtype, mlen, msg)
	}
}

// go run main.go

/*
0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string

	type	 | length  | data
*/
func createMessage(mtype int, data string) []byte {
	buf := make([]byte, 4+4+len(data))
	binary.LittleEndian.PutUint32(buf[0:], uint32(mtype))     // 0 dan ıtıbaren
	binary.LittleEndian.PutUint32(buf[4:], uint32(len(data))) //4 en ıtıbaren
	copy(buf[8:], []byte(data))                               //8 den ıtıbaren
	return buf
}

func readMessage(data []byte) (mtype, mlen uint32, msg string) {
	mtype = binary.LittleEndian.Uint32(data[0:])
	mlen = binary.LittleEndian.Uint32(data[4:])
	msg = string(data[8:])
	return mtype, mlen, msg
}
