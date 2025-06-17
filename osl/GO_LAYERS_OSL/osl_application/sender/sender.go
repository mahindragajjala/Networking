// client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	message := "Hello from Client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Message sent to server.")
}
