package main

import (
	"log"
	"net"
)

func main() {
	raddr, err := net.ResolveSCTPAddr("sctp", &net.SCTPAddr{
		IPAddrs: []net.IPAddr{{IP: net.ParseIP("127.0.0.1")}},
		Port:    9000,
	})
	if err != nil {
		log.Fatal("Resolve error:", err)
	}

	conn, err := net.DialSCTP("sctp", nil, raddr)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	message := "Hello from SCTP client"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	log.Println("Sent:", message)
}
