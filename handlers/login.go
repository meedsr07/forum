package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"forum/database"

	"golang.org/x/crypto/bcrypt"
)

// This function creates a random secret code (like: "aB3dE5...")
func generateSessionToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")

	//  If we found a cookie AND it is not empty, the user is already logged in
	if err == nil && cookie.Value != "" {

		// Redirect the user to the Home page ("/")
		http.Redirect(w, r, "/", http.StatusSeeOther)

		// STOP here! Do not run the rest of the code (Do not show the login page)
		return
	}
	// 1. If the user just wants to see the page (GET request)
	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "login.html", nil)
		return
	}

	// 2. If the user submitted the login form (POSST request)
	if r.Method == http.MethodPost {
		// a. Get the data entered by the user
		identifier := r.FormValue("identify") // Can be either username or email
		password := r.FormValue("password")

		var dbID int
		var dbPasswordHash string

		// b. Check if this user exists in the database
		err := database.DB.QueryRow("SELECT id, password FROM users WHERE email = ? OR username = ?", identifier, identifier).Scan(&dbID, &dbPasswordHash)

		if err != nil {
			// If user is not found, return an error
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
				tmpl.ExecuteTemplate(w, "login.html", map[string]interface{}{
					"Error": "Invalid email/username or password.",
				})
				return
			}
			// Database error
			ErrorHandler(w, "Server error", http.StatusInternalServerError)
			return
		}

		// c. User found! Now compare the submitted password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(dbPasswordHash), []byte(password))
		if err != nil {
			// Incorrect password
			w.WriteHeader(http.StatusUnauthorized)
			tmpl.ExecuteTemplate(w, "login.html", map[string]interface{}{
				"Error": "Invalid email/username or password.",
			})
			return
		}

		// delete old sessions
		deleteQuery := "DELETE FROM user_sessions WHERE user_id = ?"
		_, err = database.DB.Exec(deleteQuery, dbID)
		if err != nil {
			log.Println("Error deleting old sessions:", err)
		}
		// d. Credentials are correct! Create a Session Token
		sessionToken := generateSessionToken()

		// e. Save the token in the 'user_sessions' table in our Database
		_, err = database.DB.Exec("INSERT INTO user_sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", dbID, sessionToken, time.Now().Add(24*time.Hour))
		if err != nil {
			log.Println("Error saving session to DB:", err)
			ErrorHandler(w, "Error creating session, try again later", http.StatusInternalServerError)
			return
		}
		// f. Create a cookie and attach the token
		cookie := http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)

		// g. Redirect the user to the home page (Logged in!)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
