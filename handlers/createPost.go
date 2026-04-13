package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"forum/database"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}

	// Must be logged in
	userID, err := GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		ErrorHandler(w, "Bad Request", http.StatusBadRequest)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))
	categoryStr := r.FormValue("category")

	// Validate fields
	if title == "" || content == "" || categoryStr == "" {
		ErrorHandler(w, "Title, content, and category are required", http.StatusBadRequest)
		return
	}

	// Convert category to int
	categoryID, err := strconv.Atoi(categoryStr)
	if err != nil || categoryID < 1 {
		ErrorHandler(w, "Invalid category", http.StatusBadRequest)
		return
	}

	// Verify category exists
	exists, err := database.CategoryExists(categoryID)
	if err != nil || !exists {
		ErrorHandler(w, "Category not found", http.StatusBadRequest)
		return
	}

	// Insert the post
	_, err = database.DB.Exec(
		`INSERT INTO posts (user_id, title, content, category_id) VALUES (?, ?, ?, ?)`,
		userID, title, content, categoryID,
	)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
