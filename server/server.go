package main

import (
	"log"
	"net"
)

type Server struct {
	listener net.Listener
}

func newServer(addr string) *Server {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{listener: listener}
}

func (s *Server) run() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := newClient(conn)
		go client.run()
	}
}
