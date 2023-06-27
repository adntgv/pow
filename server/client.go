package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"

	powlib "github.com/adntgv/powlib/pow"
)

type Client struct {
	conn net.Conn
}

func newClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) run() {
	defer c.conn.Close()

	challenge := generateChallenge()
	c.send(challenge)
	nonce := c.getNonce()

	// Validate the PoW
	if !c.isValidPoW(challenge, nonce) {
		log.Printf("Invalid PoW nonce %s for challenge %s", nonce, challenge)
		fmt.Fprintf(c.conn, "Did not solve the challenge\n")
		return
	}

	// Send the quote to the client
	quote := getRandomQuote()

	c.send(quote)
}

func generateChallenge() string {
	return strconv.Itoa(rand.Intn(100000))
}

func (c *Client) getNonce() string {
	// Read the response from the client
	nonce, err := bufio.NewReader(c.conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	nonce = strings.TrimSpace(nonce)
	return nonce
}

func (c *Client) isValidPoW(challenge, nonce string) bool {
	return powlib.IsValidPoW(challenge, nonce)
}

func getRandomQuote() string {
	quites := []string{
		"Be yourself; everyone else is already taken.",
		"Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
		"So many books, so little time.",
		"A room without books is like a body without a soul.",
		"You only live once, but if you do it right, once is enough.",
		"Be the change that you wish to see in the world.",
		"In three words I can sum up everything I've learned about life: it goes on.",
	}

	return quites[rand.Intn(len(quites))]
}

func (c *Client) send(data string) {
	// Send the data to the client
	fmt.Fprintf(c.conn, "%s\n", data)
}
