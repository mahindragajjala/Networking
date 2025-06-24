// forward_proxy.go
package main

import (
    "io"
    "log"
    "net/http"
    "net/url"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
    target := r.URL.Query().Get("url")
    if target == "" {
        http.Error(w, "Missing URL param", http.StatusBadRequest)
        return
    }

    parsedURL, err := url.Parse(target)
    if err != nil {
        http.Error(w, "Invalid URL", http.StatusBadRequest)
        return
    }

    // Forward the request to the actual server
    resp, err := http.Get(parsedURL.String())
    if err != nil {
        http.Error(w, "Failed to fetch URL", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    // Copy response headers
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }

    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

func main() {
    http.HandleFunc("/proxy", handleProxy)
    log.Println("Forward proxy running on :8080")
    http.ListenAndServe(":8080", nil)
}

/* 
🚀 How to Use It
Run:
        go run forward_proxy.go

Call it like this from a client/browser:
        http://localhost:8080/proxy?url=http://example.com
        
This will:
        Fetch http://example.com
        Return the HTML content to the client through the proxy
*/
