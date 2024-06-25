package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

// Server struct
type Server struct {
	listenAddr string        // server listen address
	ln         net.Listener  // server listener
	quitch     chan struct{} // server quit channel
	msgch      chan Message  // server message channel
}

// NewServer creates a new server
func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}

// Start starts the server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr) // create server listener
	if err != nil {
		return err
	}
	defer ln.Close() // close server listener
	s.ln = ln        // save server listener

	go s.acceptLoop() // start server accept loop

	<-s.quitch // wait for server quit
	close(s.msgch)

	return nil // server quited
}

// acceptLoop accepts new connections
func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept() // accept new connection
		if err != nil {
			fmt.Print("accept error:", err) // print error
			continue
		}
		fmt.Println("new connection to the server:", conn.RemoteAddr()) // print new connection
		go s.readLoop(conn)                                             // start read loop for the new connection
	}
}

// readLoop reads messages from the connection

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()        // close connection
	buf := make([]byte, 2048) // buffer for reading
	for {
		n, err := conn.Read(buf) // read from connection
		if err != nil {
			fmt.Print("read error:", err) // print error
			continue
		}
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}

		conn.Write([]byte("thank you for your message!"))

	}
}

// readLoop reads messages from the connection line by line
//func (s *Server) readLoop(conn net.Conn) {
//	defer conn.Close() // close connection
//
//	scanner := bufio.NewScanner(conn)
//	for scanner.Scan() {
//		msg := scanner.Bytes() // get bytes from scanner
//		s.msgch <- Message{
//			from:    conn.RemoteAddr().String(),
//			payload: msg,
//		}
//	}
//
//	if err := scanner.Err(); err != nil {
//		fmt.Println("read error:", err) // print error
//	}
//}

func main() {
	server := NewServer(":3000") // create server

	go func() {
		for msg := range server.msgch {
			fmt.Printf("received message from connection (%s):%s\n", msg.from, string(msg.payload))
		}
	}()
	log.Fatal(server.Start()) // start server
}
