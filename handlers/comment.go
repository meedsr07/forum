package handlers

import (
	// Adjust import to match your module name
	"database/sql"
	"forum/models"
	"net/http"
	"strconv"
)

func CreateCommentHandler(db *sql.DB) http.HandlerFunc {
	// 2. Return the actual handler function
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST requests (form submissions)
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the form data
		r.ParseForm()
		content := r.FormValue("content")
		postIDStr := r.FormValue("post_id")

		postID, err := strconv.Atoi(postIDStr)
		if err != nil || content == "" {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// GET LOGGED IN USER (Hardcoded to 1 for now until Auth is done)
		userID := 1
		// LATER: userID := GetUserFromSession(r)

		// 3. FIX IS HERE: Use the 'db' passed into the parent function!
		err = models.CreateComment(db, postID, userID, content)
		if err != nil {
			http.Error(w, "Failed to create comment", http.StatusInternalServerError)
			return
		}

		// Redirect the user back to the post page
		http.Redirect(w, r, "/post?id="+postIDStr, http.StatusSeeOther)
	}
}

// --------------------------
// maybe modify, delete later
//---------------------------