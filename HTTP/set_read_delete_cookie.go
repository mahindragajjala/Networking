package main

import (
	"fmt"
	"net/http"
)

/*
🍪 1. What is a Cookie? — Theory
A cookie is a small piece of data stored in the browser (client) by the server via HTTP headers.

🔑 Structure of a Cookie (Key Attributes)
Attribute	Description
Name	The name of the cookie (e.g., user_id)
Value	The value stored (e.g., 12345)
Path	URL path where the cookie is valid (e.g., /user)
Domain	Domain for which the cookie is valid (e.g., .example.com)
Secure	Cookie is sent only over HTTPS
HttpOnly	Can't be accessed by JavaScript → protects against XSS
MaxAge	Time in seconds until the cookie expires
Expires	Exact date/time when the cookie expires (alternative to MaxAge)
SameSite	Prevents CSRF (cross-site request forgery) attacks
*/
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

