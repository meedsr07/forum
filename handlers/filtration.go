package handlers

import (
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
	userID, sessionErr := GetUserID(r)
	filter := r.URL.Query().Get("filter")

	var Post []models.Post
	var err error

	switch filter {

	case "myposts":
		if sessionErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		Post, err = database.GetMyPosts(database.DB, userID)
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}

	case "liked":
		if sessionErr != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		Post, err = database.GetLikedPosts(database.DB, userID)
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}

	default:
		Post, err = database.Getallpost(database.DB)
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "page not found", 404)
		return
	}
	tmpl.Execute(w, Post)
}