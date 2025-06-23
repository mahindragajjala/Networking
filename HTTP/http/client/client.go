// client.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Step 1: Dial TCP manually
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("TCP Dial error:", err)
	}
	defer conn.Close()
	fmt.Println("TCP Client: Connected to server")

	// Step 2: Send HTTP request over TCP
	request := "GET / HTTP/1.1\r\n" +
		"Host: localhost\r\n" +
		"Connection: close\r\n\r\n"

	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	fmt.Println("TCP Client: Sent HTTP request")

	// Step 3: Read HTTP response
	resp := bufio.NewReader(conn)
	fmt.Println("===== HTTP RESPONSE FROM SERVER =====")
	for {
		line, err := resp.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}

