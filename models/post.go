package models

import (
	"forum/database"
	"forum/handlers"
	"net/http"
	"strings"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handlers.ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var userID int
	err = database.DB.QueryRow("SELECT user_id FROM sessions WHERE id = ?", cookie.Value).Scan(&userID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}
	err = r.ParseForm()
	if err != nil {

		handlers.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)
	category = strings.TrimSpace(category)
	if title == "" || content == "" || category == "" || len(title) > 200 || len(content) > 4096 {
		handlers.ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var categoryID int
	err = database.DB.QueryRow(
		"SELECT id FROM categories WHERE name = ?",
		category,
	).Scan(&categoryID)

	if err != nil {
		handlers.ErrorHandler(w, "Invalid category", http.StatusBadRequest)
		return
	}
	_, err = database.DB.Exec(
		`INSERT INTO posts (user_id, title, content, category_id)
	 VALUES (?, ?, ?, ?)`,
		userID, title, content, categoryID,
	)
	if err != nil {
		handlers.ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
