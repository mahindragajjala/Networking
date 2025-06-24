
/* 
A frontend proxy receives client requests.
It forwards requests to one or more backend services.
It hides internal server details, adds SSL, caching, etc. 

🧠 How It Works
Client talks to proxy (e.g., proxy.myapp.com)
Proxy forwards to internal services like localhost:8081
*/
//Backend Go Server
// backend.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Backend Service!")
    })

    log.Println("Backend server running on :8081")
    http.ListenAndServe(":8081", nil)
}


//Reverse Proxy Server
// reverse_proxy.go
package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    target, _ := url.Parse("http://localhost:8081")
    proxy := httputil.NewSingleHostReverseProxy(target)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        proxy.ServeHTTP(w, r)
    })

    log.Println("Reverse Proxy running on :8080")
    http.ListenAndServe(":8080", nil)
}

