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

	_, err = conn.Write([]byte("hello!!!"))
	if err != nil {
		fmt.Println("write error:", err)
	}

}
