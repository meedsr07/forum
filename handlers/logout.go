package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/database"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get the cookie from the user's browser
	cookie, err := r.Cookie("session_token")
	if err == nil {
		// 2. Delete the session from the database
		database.DB.Exec("DELETE FROM user_sessions WHERE session_token = ?", cookie.Value)
	}

	// 3. Delete the cookie from the browser (by setting expiration time in the past)
	deletedCookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expired 1 hour ago
		HttpOnly: true,
		Path:     "/",
	}
	
	http.SetCookie(w, &deletedCookie)

	// 4. Send the user back to the homepage or login page
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
