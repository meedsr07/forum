package handlers

import (
	"bytes"
	"fmt"
	"html/template"
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
		fmt.Println("error ")
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
<<<<<<< HEAD
		Post, err = database.GetMyPosts(userID)
=======
		posts, err = database.GetMyPosts(database.DB, userID)
>>>>>>> test-merge
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}

	case "liked":
		if sessionErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
<<<<<<< HEAD
		Post, err = database.GetLikedPosts(userID)
=======
		posts, err = database.GetLikedPosts(database.DB, userID)
>>>>>>> test-merge
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}

	default:
<<<<<<< HEAD
		Post, err = database.Getallpost()
=======
		posts, err = database.Getallpost(database.DB)
>>>>>>> test-merge
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

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "page not found", 404)
		return
	}
<<<<<<< HEAD
	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, Post); err != nil {
		// If the template cannot be executed, return a generic 500 error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	w.Write(buff.Bytes())
=======
	tmpl.Execute(w, pageData)
>>>>>>> test-merge
}
