package handlers

import (
	"net/http"
	"strconv"

	"forum/database"
)

func Likehandler(w http.ResponseWriter, r *http.Request) {
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

	postIDStr := r.FormValue("post_id")
	valueStr := r.FormValue("value")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil || postID < 1 {
		ErrorHandler(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil || (value != 1 && value != -1) {
		ErrorHandler(w, "Invalid vote value", http.StatusBadRequest)
		return
	}

	if err := database.HandleVote(userID, postID, value); err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}
