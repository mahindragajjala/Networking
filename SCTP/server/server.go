package main

import (
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveSCTPAddr("sctp", &net.SCTPAddr{
		IPAddrs: []net.IPAddr{{IP: net.ParseIP("127.0.0.1")}},
		Port:    9000,
	})
	if err != nil {
		log.Fatal("Resolve error:", err)
	}

	listener, err := net.ListenSCTP("sctp", addr)
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	defer listener.Close()
	log.Println("SCTP Server listening on", addr)

	for {
		conn, err := listener.AcceptSCTP()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		log.Println("Connection accepted:", conn.RemoteAddr())

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Read error:", err)
			conn.Close()
			continue
		}
		log.Println("Received:", string(buf[:n]))
		conn.Close()
	}
}
