package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"forum/database"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(405), 405)
		return
	}

	r.ParseForm()
	content := strings.TrimSpace(r.FormValue("content"))
	postIDStr := r.FormValue("post_id")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		ErrorHandler(w, http.StatusText(400), 400)
		return
	}

	userID, err := GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/login?next=/post/"+postIDStr, http.StatusSeeOther)
		return
	}

	if content == "" {
		http.Redirect(w, r, "/post/"+postIDStr+"?error=Comment cannot be empty.", http.StatusSeeOther)
		return
	}

	if len(content) > 2048 {
		http.Redirect(w, r, "/post/"+postIDStr+"?error=Comment is too long (max 2048 characters).", http.StatusSeeOther)
		return
	}

	for strings.Contains(content, "\r\n\r\n") {
		content = strings.ReplaceAll(content, "\r\n\r\n", "\r\n")
	}

	err = database.CreateComment(postID, userID, content)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}
