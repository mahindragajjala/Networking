package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    address := "localhost:38412"

    conn, err := net.Dial("tcp", address)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    message := "NGAP: InitialUEMessage [Mock]"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }

    fmt.Println("Sent NGAP message to server")
    time.Sleep(1 * time.Second)
}
