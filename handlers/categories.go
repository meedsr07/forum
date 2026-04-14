package handlers

import (
	"net/http"
	"strconv"

	"forum/database"
	"forum/models"
)

func FilterPostsHandler(w http.ResponseWriter, r *http.Request ,IdCategory string) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}

	catStr := IdCategory
	if catStr == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	categoryID, err := strconv.Atoi(catStr)
	if err != nil || categoryID < 1 {
		ErrorHandler(w, "Invalid category", http.StatusBadRequest)
		return
	}

	posts, err := database.GetPostsByCategory(categoryID)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// Build auth state
	userID, sessionErr := GetUserID(r)
	pageData := models.PageData{
		Posts: posts,
	}
	pageData.Category, err = database.GetAllCategories()
	if err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return 
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
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
}
