package main

import (
	"fmt"
	"net/http"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Create the cookie struct
	cookie := http.Cookie{
		Name:     "username",       // Key
		Value:    "mahindra",       // Value
		Path:     "/",              // Valid for all paths
		HttpOnly: true,             // Not accessible from JavaScript
		Secure:   false,            // Set to true in HTTPS
		MaxAge:   3600,             // 1 hour (in seconds)
		// Expires: time.Now().Add(1 * time.Hour), // alternative to MaxAge
	}

	// Add Set-Cookie to the response header
	http.SetCookie(w, &cookie)

	// Respond to the client
	fmt.Fprintln(w, "✅ Cookie set successfully!")
}

func main() {
	http.HandleFunc("/set", setCookieHandler)
	fmt.Println("Server running at http://localhost:8080/set")
	http.ListenAndServe(":8080", nil)
}

