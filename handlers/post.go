package handlers

import (
	"bytes"
	"net/http"
	"strconv"

	"forum/database"
	"forum/models"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}
	path := r.URL.Path
	if len(path) < 6 || path[:6] != "/post/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	postID, err := strconv.Atoi(path[6:])
	if err != nil || postID < 0 {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	post, err := database.GetOnePost(postID)
	if err != nil {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}

	comments, err := database.GetCommentsByPost(postID)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	Likess, Dislikess := database.GetPostVotes(postID)
	
	commentVotes := make(map[int]struct {
		Likes    int
		Dislikes int
	})

	for _, c := range comments {
		likes, dislikes := database.GetCommentVotes(c.ID)

		commentVotes[c.ID] = struct {
			Likes    int
			Dislikes int
		}{
			Likes:    likes,
			Dislikes: dislikes,
		}
	}
	userID, sessionErr := GetUserID(r)
	data := models.PostPageData{
		Post:     post,
		Comments: comments,
		Likes:    Likess,
		Dislikes: Dislikess,
	}
	if sessionErr == nil {
		var username string
		dbErr := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
		if dbErr == nil {
			data.IsLoggedIn = true
			data.Username = username
		}
	}

	var buff bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buff, "post.html", data); err != nil {
		ErrorHandler(w, "intenal srever error", 500)
		return
	}
	w.Write(buff.Bytes())
}
