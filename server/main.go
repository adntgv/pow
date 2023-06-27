package main

import (
	"log"
)

const (
	host = "0.0.0.0"
	port = "8000"
)

func main() {
	server := newServer(host + ":" + port)
	defer server.listener.Close()

	log.Println("Server is listening on", host+":"+port)
	server.run()
}
