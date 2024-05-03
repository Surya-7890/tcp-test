package server

import (
	"fmt"
	"net"
)

type Server struct {
	address  string
	listener net.Listener
	exit     chan struct{}
	Msg      chan []byte
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		exit:    make(chan struct{}),
		Msg:     make(chan []byte, 2),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	s.listener = listener
	defer s.listener.Close()
	go s.acceptLoop()
	<-s.exit
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("error while acception connection", err)
			continue
		}
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error while reading message")
			continue
		}
		s.Msg <- buffer[:length]
	}
}
