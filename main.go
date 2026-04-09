package main

import (
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> origin/main
	"log"
	"net/http"

	"forum/database"
<<<<<<< HEAD
)

func main() {
	// Init DB + create all tables
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal("could not init db:", err)
	}
	defer db.Close()

	// Populate with test data (remove in production)
	if err := database.SeedDB(db); err != nil {
		log.Fatal("could not seed db:", err)
	}

	log.Println("Forum running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
=======

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

	http.HandleFunc("/static/", handlers.StaticHandlers)
	fmt.Println("server is start in http://localhost:8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}
>>>>>>> origin/main
}
