package handlers

import (
	"bytes"
	"fmt"
	"forum/database"
	"forum/models"
	"html/template"
	"net/http"
	"strconv"
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

	data := models.PostPageData{
		Post:     post,
		Comments: comments,
	}

	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, data); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(buff.Bytes())
}
