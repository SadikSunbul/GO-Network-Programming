package main

import (
	"fmt"
	"net"
)

func main() {
	ls, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}
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
		buff := make([]byte, 8)
		_, err := conn.Read(buff[:])
		if err != nil {
			fmt.Println("read error:", err)
			conn.Close()
			break
		}
		fmt.Sprintf("mssage client : %s\n", buff)

		_, err = conn.Write(buff)
		if err != nil {
			fmt.Println("write error:", err)
			continue
		}

	}
}
