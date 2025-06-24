package main

import (
	"fmt"
	"net/http"
)

func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Read a cookie by name
	cookie, err := http.Cookie("username")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "❌ Cookie not found", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print cookie value
	fmt.Fprintf(w, "✅ Hello, %s!", cookie.Value)
}

func main() {
	http.HandleFunc("/read", readCookieHandler)
	http.ListenAndServe(":8080", nil)
}
