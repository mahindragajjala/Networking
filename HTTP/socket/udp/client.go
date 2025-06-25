// udp_client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := net.UDPAddr{
		Port: 9001,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send message
	message := "Hello from UDP client"
	conn.Write([]byte(message))

	// Receive reply
	buffer := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buffer)
	fmt.Printf("Server replied: %s\n", string(buffer[:n]))
}
