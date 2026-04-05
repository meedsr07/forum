package handlers

import (
	"database/sql"
	"log"
	"net/http"
)

// HomeHandler handles requests to the main forum page
func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// 1. Check if user is logged in (but don't force it!)
	userID, err := GetUserIDFromCookie(r, db)

	isLoggedIn := (err == nil)
	username := ""
	if r.URL.Path != "/homepage" && r.URL.Path != "/" {
		ErrorHandler(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	// 2. If logged in, get the real username from the database
	if isLoggedIn {
		err = db.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
		if err != nil {
			log.Println("Error fetching username:", err)
			ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	// 3. Prepare the data for HTML
	data := map[string]interface{}{
		"IsLoggedIn": isLoggedIn,
		"Username":   username,
	}

	// 4. Prevent browser caching (So Logout works perfectly)
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// 5. Render the page
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error rendering home page:", err)
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
