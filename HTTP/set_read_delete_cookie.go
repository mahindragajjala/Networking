package main

import (
	"fmt"
	"net/http"
)
/*
🔍 What’s happening?
Server sends Set-Cookie: username=mahindra; Path=/; HttpOnly

Browser stores it and includes it in every future request to /
*/
func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Create a cookie struct
	cookie := http.Cookie{
		Name:     "username",
		Value:    "mahindra",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,    // Set to true in HTTPS only
		MaxAge:   3600,     // 1 hour
	}

	// Send it to the client
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie set successfully!")
}
/*
When client sends request with cookie → server reads it from r.Cookie("key")

Example Request Header: Cookie: username=mahindra
*/
func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Read a specific cookie
	cookie, err := r.Cookie("username")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", cookie.Value)
}
//delete the cookie
func deleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Set the MaxAge to -1 to delete the cookie
	cookie := http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie deleted!")
}

func main() {
	http.HandleFunc("/set", setCookieHandler)
	http.HandleFunc("/read", readCookieHandler)
	http.HandleFunc("/delete", deleteCookieHandler)

	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

