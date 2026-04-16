package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"

	"forum/database"

	"golang.org/x/crypto/bcrypt"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{4,20}$`)
	emailRegex    = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")

	if err == nil && cookie != nil && cookie.Value != "" {
		var dbToken string
		var expiresAt time.Time
		errDB := database.DB.QueryRow("SELECT session_token, expires_at FROM user_sessions WHERE session_token = ?", cookie.Value).Scan(&dbToken, &expiresAt)

		if errDB == nil {
			if time.Now().Before(expiresAt) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				database.DB.Exec("DELETE FROM user_sessions WHERE session_token = ?", cookie.Value)
			}
		}
	}
	// 1. If GET request: Show the register page
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		tmpl.ExecuteTemplate(w, "register.html", nil)
		return
	}

	// 2. If POST request: Save the new user
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Error reading form",
			})
			return
		}
		///////Email & Username check\\\\\\
		username := r.FormValue("username")
		email := r.FormValue("email")

		if !usernameRegex.MatchString(username) {
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Username must contain only letters and numbers.",
			})
			return
		}

		if !emailRegex.MatchString(email) {
			w.WriteHeader(http.StatusBadRequest)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Please enter a valid email address.",
			})
			return
		}

		var existingID int
		err = database.DB.QueryRow("SELECT id FROM users WHERE email = ? OR username = ?", email, username).Scan(&existingID)

		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusConflict)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Email or Username already exists",
			})
			return
		}
		///////Password check\\\\
		password := r.FormValue("password")
		// 3. Hash the password before saving it
		if len(password) > 72 {
			w.WriteHeader(http.StatusBadRequest) // 400 Bad Request

			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Password is too long! (72 characters max)",
			})

			return
		}
		if len(password) < 8 {
			w.WriteHeader(http.StatusBadRequest) // 400 Bad Request

			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Password is too short! (8 characters minimum)",
			})

			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Bcrypt hash error:", err)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Error creating account, try again later",
			})
			return
		}

		_, err = database.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, string(hashedPassword))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Database insert error:", err)
			tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
				"Error": "Could not create account",
			})
			return
		}

		// Success! Redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// 3. Reject anything else
	ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
