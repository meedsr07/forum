package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// 1. If GET request: Show the register page
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, "Error loading page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	// 2. If POST request: Save the new user
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error reading form", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if username == "" || email == "" || password == "" {
			http.Error(w, "Please fill all fields", http.StatusBadRequest)
			return
		}

		var existingID int
		err = db.QueryRow("SELECT id FROM users WHERE email = ? OR username = ?", email, username).Scan(&existingID)

		if err != sql.ErrNoRows {
			http.Error(w, "Email or Username already exists", http.StatusConflict)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, string(hashedPassword))
		if err != nil {
			log.Println("Database insert error:", err)
			http.Error(w, "Could not create account", http.StatusInternalServerError)
			return
		}

		// Success! Redirect to login page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// 3. Reject anything else
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
