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
	path := r.URL.Path
	if len(path) < 6 || path[:6] != "/like/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	postID, err := strconv.Atoi(path[6:])
	if err != nil || postID < 0 {
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

	if err := database.HandleVote(userID, postID, 1); err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	postIDStr := strconv.Itoa(postID)
	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}

func Likecommenthandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}
	path := r.URL.Path
	if len(path) < 14 || path[:14] != "/comment/like/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	commentID, err := strconv.Atoi(path[14:])
	if err != nil || commentID < 0 {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}
		postIDStr := r.URL.Query().Get("PostId")

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

	if err := database.HandleVotecomment(userID, commentID, 1); err != nil {
		ErrorHandler(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}
