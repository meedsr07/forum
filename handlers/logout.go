package handlers

import (
	"database/sql"
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// 1. Get the cookie from the user's browser
	cookie, err := r.Cookie("session_token")
	if err == nil {
		// 2. Delete the session from the database
		db.Exec("DELETE FROM sessions WHERE session_token = ?", cookie.Value)
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
