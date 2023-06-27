package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	powlib "github.com/adntgv/powlib/pow"
)

type Client struct {
	conn net.Conn
}

func newClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	return &Client{conn: conn}, fmt.Errorf("could not connect to server: %v", err)
}

func (c *Client) getChallange() (string, error) {
	// Read the challenge from the server
	challenge, err := c.receive()
	if err != nil {
		return "", fmt.Errorf("could not receive challenge: %v", err)
	}

	log.Println("Received challenge:", challenge)
	return challenge, nil
}

func (c *Client) solveChallenge(challenge string) string {
	// Solve the PoW challenge
	nonce := powlib.SolveChallenge(challenge)

	return nonce
}

func (c *Client) getQuote(nonce string) (string, error) {
	// Send the nonce to the server
	c.send(nonce)

	return c.receive()
}

func (c *Client) send(msg string) error {
	_, err := fmt.Fprintf(c.conn, "%s\n", msg)
	return err
}

func (c *Client) receive() (string, error) {
	// Read the response from the server
	// abort if timeout
	ch := make(chan string)
	che := make(chan error)

	go func(conn net.Conn) {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			che <- fmt.Errorf("could not receive message: %v", err)
			return
		}
		msg = strings.TrimSpace(msg)
		ch <- msg
	}(c.conn)

	select {
	case err := <-che:
		return "", err
	case <-time.After(5 * time.Second):
		return "", fmt.Errorf("Timeout")
	case msg := <-ch:
		return msg, nil
	}
}