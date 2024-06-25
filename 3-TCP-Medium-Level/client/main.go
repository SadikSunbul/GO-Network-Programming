package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp4", ":3000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go func() {
		for {
			buff := make([]byte, 8)
			_, err := conn.Read(buff[:])
			_, _, msg := readMessage(buff)
			if err != nil {
				fmt.Println("read error:", err)
				conn.Close()
				break
			}
			fmt.Printf("mssage from server : %s \n", msg)
		}
	}()

	for i := 0; i < 10; i++ {
		data := createMessage(MessageTypeText, "hello from client")
		_, err = conn.Write(data)
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
