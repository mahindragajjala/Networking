package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var sessionStore = make(map[string]string) // sessionID -> username

// Utility to generate random session ID
func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// Middleware to check session
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		if username, ok := sessionStore[cookie.Value]; ok {
			c.Set("username", username)
			c.Next()
		} else {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()

	// GET /login (login form)
	r.GET("/login", func(c *gin.Context) {
		c.Writer.WriteString(`
			<form method="POST" action="/login">
				Username: <input name="username" />
				Password: <input name="password" />
				<button type="submit">Login</button>
			</form>
		`)
	})

	// POST /login (form submission)
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "mahindra" && password == "1234" {
			sessionID := generateSessionID()
			sessionStore[sessionID] = username

			c.SetCookie("session_id", sessionID, 300, "/", "", false, true)
			c.Redirect(http.StatusSeeOther, "/dashboard")
		} else {
			c.String(http.StatusUnauthorized, "Invalid credentials")
		}
	})

	// GET /dashboard (protected)
	r.GET("/dashboard", authRequired(), func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.String(http.StatusOK, "Welcome to your dashboard, %s!\n\n<a href='/logout'>Logout</a>", username)
	})

	// GET /logout
	r.GET("/logout", func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		if err == nil {
			delete(sessionStore, cookie.Value)
			c.SetCookie("session_id", "", -1, "/", "", false, true)
		}
		c.Redirect(http.StatusSeeOther, "/login")
	})

	// Root redirect
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/login")
	})

	fmt.Println("Server running at http://localhost:8080")
	r.Run(":8080")
}

/* 
| Operation | Route        | Method   | Description                        |
| --------- | ------------ | -------- | ---------------------------------- |
| Login     | `/login`     | GET/POST | Set session in cookie & memory     |
| Dashboard | `/dashboard` | GET      | Protected route using middleware   |
| Logout    | `/logout`    | GET      | Clear session from memory & cookie |
*/

//SEE THE MIDDLEWARE 
/*
| Middleware Type      | Purpose                                    | Example      |
| Authentication   | Check session/cookie/token                 | `/dashboard` |
| Logging          | Log all requests (method, URL, duration)   | Global       |
| Rate Limiting    | Block users making too many requests       | `/api/`      |
| Recovery         | Catch and handle panics                    | Global       |
| CORS             | Add headers to allow cross-origin requests | Global       |
| Input Validation | Check for valid headers, request bodies    | `/api/post`  |

*/
