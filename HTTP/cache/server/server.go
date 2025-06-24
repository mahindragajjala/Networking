package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	etagValue = `"v1"` // fake ETag
	data      = []byte("This is cached data!")
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Check for ETag match
	if match := r.Header.Get("If-None-Match"); match == etagValue {
		w.Header().Set("ETag", etagValue)
		w.WriteHeader(http.StatusNotModified) // 304
		fmt.Println("✅ ETag matched, returning 304 Not Modified")
		return
	}

	// Send full response with caching headers
	w.Header().Set("Cache-Control", "max-age=300")
	w.Header().Set("ETag", etagValue)
	w.Header().Set("Expires", time.Now().Add(5*time.Minute).Format(http.TimeFormat))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	fmt.Println("📦 Sent 200 OK with full content")
}

func main() {
	http.HandleFunc("/data", handler)
	fmt.Println("🚀 Server running at http://localhost:8080/data")
	http.ListenAndServe(":8080", nil)
}
