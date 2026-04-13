package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"forum/database"
)

// LogoutHandler handles the user logout process
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	// 1. Get the session cookie from the user's browser
	cookie, err := r.Cookie("session_token")

	// 2. If the cookie exists, delete the session from the database
	if err == nil && cookie.Value != "" {

		query := "DELETE FROM user_sessions WHERE session_token = ?"
		_, err := database.DB.Exec(query, cookie.Value)

		if err != nil {
			log.Println("Error deleting sessijhon from DB:", err)
		}
	}

	// 3. Create a "dead" cookie to delete the old one in the browser
	deletedCookie := &http.Cookie{
		Name:     "session_token",
		Value:    "",   // Empty value
		Path:     "/",  // Must match the original cookie path
		MaxAge:   -1,   // -1 tells the browser to delete it immediately!
		HttpOnly: true, // Keeps it secure from JavaScript
	}

	// 4. Send this "dead" cookie to the user's browser
	http.SetCookie(w, deletedCookie)

	// 5. Redirect the user back to the Home page ("/")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// GetUserIDFromCookie checks the cookie and returns the logged-in User's ID.
func GetUserIDFromCookie(r *http.Request) (int, error) {
	// 1. Does the user have a session cookie?
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, fmt.Errorf("no cookie found")
	}

	var userID int
	var expiresAt time.Time

	// 2. Does this token exist in our database?
	err = database.DB.QueryRow("SELECT user_id, expires_at FROM user_sessions WHERE session_token = ?", cookie.Value).Scan(&userID, &expiresAt)
	if err != nil {
		return 0, fmt.Errorf("invalid session token")
	}

	// 3. Is the session expired?
	if time.Now().After(expiresAt) {
		return 0, fmt.Errorf("session expired")
	}

	// 4. Success! Return the user's ID
	return userID, nil
}
