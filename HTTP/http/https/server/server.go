// server.go
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/secure", func(c *gin.Context) {
        fmt.Println("🔒 Secure request from:", c.ClientIP())
        c.JSON(200, gin.H{"message": "Hello over HTTPS"})
    })

    fmt.Println("🔐 Running HTTPS server at https://localhost:8443")
    // Run HTTPS server with cert and key
    err := r.RunTLS(":8443", "cert.pem", "key.pem")
    if err != nil {
        panic(err)
    }
}

