package main

import (
    "fmt"
    "net"
)

func main() {
    address := ":38412" // NGAP uses SCTP on port 38412, here using TCP for simplicity

    listener, err := net.Listen("tcp", address)
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("NGAP Server is listening on", address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Failed to accept connection:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }
    fmt.Println("Received NGAP message:", string(buffer[:n]))
}
