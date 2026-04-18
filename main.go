package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/database"

	"forum/handlers"
)

func main() {
	database.InitializeDB()
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/Post/CreatePost", handlers.CreateNewPost)
	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/comment/create", handlers.CreateCommentHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/like/", handlers.Likehandler)
	http.HandleFunc("/dislike/", handlers.DisLikehandler)
	http.HandleFunc("/comment/like/", handlers.Likecommenthandler)
	http.HandleFunc("/comment/dislike/", handlers.DisLikecommenthandler)


	http.HandleFunc("/static/", handlers.StaticHandlers)
	fmt.Println("server is start in http://localhost:8089")
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal(err)
	}
}
