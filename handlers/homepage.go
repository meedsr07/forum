package handlers

import (
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Check if the user has a "session_token" cookie
	_, err := r.Cookie("session_token")

	// 2. If the cookie is missing (user is not logged in)
	if err != nil {
		// Redirect them to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// 3. If the cookie is found (user is logged in), show the home page
	err = tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println("Error executing home template:", err)
	}
}
