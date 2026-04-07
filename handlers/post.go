package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"forum/database"
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

	template, err := template.ParseFiles("templates/post.html")
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	var buff bytes.Buffer
	if err := template.Execute(&buff, post); err != nil {
		// If the template cannot be executed, return a generic 500 error
		fmt.Println(post)
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(buff.Bytes())
}
