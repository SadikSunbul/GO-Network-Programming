package main

import (
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
			if err != nil {
				fmt.Println("read error:", err)
				conn.Close()
				break
			}
			fmt.Printf("mssage from server : %s \n", buff)
		}
	}()

	for i := 0; i < 10; i++ {
		_, err = conn.Write([]byte("hello!!!"))
		if err != nil {
			fmt.Println("write error:", err)
		}
	}

	for {
		select {}
	}

}
