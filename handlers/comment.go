package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Println("post_id:", r.FormValue("post_id"))
	fmt.Println("content:", r.FormValue("content"))
	postID, err := strconv.Atoi(postIDStr)
	if err != nil || content == "" {
		ErrorHandler(w, http.StatusText(400), 400)
		return
	}

	userID, err := GetUserID(r)
	if err != nil {
		ErrorHandler(w, http.StatusText(401), 401)
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
