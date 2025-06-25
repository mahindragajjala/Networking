// tcp_client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send message
	message := "Hello from TCP client"
	conn.Write([]byte(message))

	// Read response
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Printf("Server replied: %s\n", string(buffer[:n]))
}
