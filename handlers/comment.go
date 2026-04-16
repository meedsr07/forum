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
	content := r.FormValue("content")
	postIDStr := r.FormValue("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if len(content) > 2048 {
		ErrorHandler(w,http.StatusText(405) , 405)
	}
	userID, err := GetUserID(r)
	if err != nil {
		http.Redirect(w, r, "/login?next=/post/"+postIDStr, http.StatusSeeOther)
		return
	}
	
	if err != nil || strings.TrimSpace(content) == "" {
		ErrorHandler(w, http.StatusText(400), 400)
		return
	}	

	err = database.CreateComment(postID, userID, content)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}
	// Redirect the user back to the post page
	http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
}
