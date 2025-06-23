// client.go
package main

import (
    "crypto/tls"
    "fmt"
    "io"
    "net/http"
)

func main() {
    // Custom HTTP client that ignores TLS cert check (not for production!)
    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }

    resp, err := client.Get("https://localhost:8443/secure")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println("📥 HTTPS Response:", string(body))
}

