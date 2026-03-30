package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"forum/database" // Import our database file (Make sure your module name is 'forum')

	"golang.org/x/crypto/bcrypt" // The Security package
)

// RegisterHandler processes the registration form
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow POST method (because we are sending form data)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Read the data from the HTML form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error reading form", http.StatusBadRequest)
		return
	}

	// 3. Get the values from the input fields (using the 'name' attribute from HTML)
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// 4. Check if any field is empty (Basic validation)
	if username == "" || email == "" || password == "" {
		http.Error(w, "Please fill all fields", http.StatusBadRequest)
		return
	}

	// 5. Check if the email or username is already taken in the database
	var existingID int
	// We ask SQLite: "Give me the ID of a user who has this email OR username"
	err = database.DB.QueryRow("SELECT id FROM users WHERE email = ? OR username = ?", email, username).Scan(&existingID)

	if err != sql.ErrNoRows {
		// If the error is NOT 'NoRows', it means a user was found!
		http.Error(w, "Email or Username already exists", http.StatusConflict) // 409 Conflict Error
		return
	}

	// 6. Hash the password (Make it secret and safe)
	// bcrypt.DefaultCost is 10 (a good balance of security and speed)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error while securing password", http.StatusInternalServerError)
		return
	}

	// 7. Save the new user into the SQLite database
	// We use '?' (Placeholders) to protect our DB from hackers (SQL Injection)
	_, err = database.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, string(hashedPassword))
	if err != nil {
		log.Println("Database insert error:", err)
		http.Error(w, "Could not create account", http.StatusInternalServerError)
		return
	}

	// 8. Success! Send the user back to the login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
