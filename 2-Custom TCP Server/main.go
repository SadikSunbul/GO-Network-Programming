package main

import (
	"fmt"
	"log"
	"net"
)

// Server struct
type Server struct {
	listenAddr string        // server listen address
	ln         net.Listener  // server listener
	quitch     chan struct{} // server quit channel
}

// NewServer creates a new server
func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
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
		msg := buf[:n]           // save message
		fmt.Println(string(msg)) // print message
	}
}

func main() {
	server := NewServer(":3000") // create server
	log.Fatal(server.Start())    // start server
}
