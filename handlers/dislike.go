package handlers

import (
	"net/http"
	"strconv"

	"forum/database"
)

func DisLikehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}
	path := r.URL.Path
	if len(path) < 9 || path[:9] != "/dislike/" {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	postID, err := strconv.Atoi(path[9:])
	if err != nil || postID < 0 {
		ErrorHandler(w, http.StatusText(300), 300)
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

	if err := database.HandleVote(userID, postID, -1); err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	postIDStr := strconv.Itoa(postID)
	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}


func DisLikecommenthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}
	path := r.URL.Path
	if len(path) < 17 || path[:17] != "/comment/dislike/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	commentID, err := strconv.Atoi(path[17:])
	if err != nil || commentID < 0 {
		ErrorHandler(w, http.StatusText(404), 404)
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

	if err := database.HandleVote(userID, commentID, 1); err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	commentIDStr := strconv.Itoa(commentID)
	http.Redirect(w, r, "/post/"+commentIDStr, http.StatusSeeOther)
}
