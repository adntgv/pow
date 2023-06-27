package main

import (
	"fmt"
	"log"
	"os"
)

var (
	host = "localhost"
	port = "8000"
)

func main() {
	client, err := newClient(host + ":" + port)
	if err != nil {
		log.Fatalln(err)
	}

	challenge, err := client.getChallange()
	if err != nil {
		log.Fatalln(err)
	}

	nonce := client.solveChallenge(challenge)

	quote, err := client.getQuote(nonce)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Received quote:", quote)
}

func init() {
	envHost := os.Getenv("HOST")
	if envHost != "" {
		host = envHost
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}
}
