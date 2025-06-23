// server.go
package main

import (
	//"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// Step 1: Create a TCP listener
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("TCP Listen error:", err)
	}
	fmt.Println("TCP Server listening on :8080")

	// Step 2: Create HTTP server with handler
	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("===== HTTP REQUEST RECEIVED =====")
			fmt.Println("Method:", r.Method)
			fmt.Println("URL:", r.URL)
			fmt.Println("Header:", r.Header)
			w.Write([]byte("Hello from TCP+HTTP Server"))
		}),
	}

	// Step 3: Serve using custom TCP listener
	log.Fatal(httpServer.Serve(listener))
}

