// udp_server.go
package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 9001,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on port 9001...")

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Received from %s: %s\n", clientAddr.String(), string(buffer[:n]))

		// Reply to client
		conn.WriteToUDP([]byte("Hello from UDP server"), clientAddr)
	}
}
