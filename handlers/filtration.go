package handlers

import (
	"net/http"

	"forum/database"
	"forum/models"
)

func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}
	token := cookie.Value

	var userID int
	err = database.DB.QueryRow("SELECT user_id FROM user_sessions  WHERE session_token = ?", token).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	userID, sessionErr := GetUserID(r)
	filter := r.URL.Query().Get("filter")

	var posts []models.Post
	var err error

	switch filter {
	case "myposts":
		if sessionErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetMyPosts(userID)
		if err != nil {
			ErrorHandler(w, "internal server error 1", 500)
			return
		}

	case "liked":
		if sessionErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetLikedPosts(userID)
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}

	default:
		posts, err = database.Getallpost()
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}
	}

	// Build auth state for the navbar
	pageData := models.PageData{
		Posts: posts,
	}
	if sessionErr == nil {
		var username string
		dbErr := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
		if dbErr == nil {
			pageData.IsLoggedIn = true
			pageData.Username = username
		}
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", pageData); err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}
	tmpl.Execute(w, pageData)
}
