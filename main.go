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
	http.HandleFunc("", handlers.Likehandler)

	http.HandleFunc("/static/", handlers.StaticHandlers)
	fmt.Println("server is start in http://localhost:8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}
}
