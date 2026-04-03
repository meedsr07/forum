package main

import (
	"fmt"
	"forum/database"
	"forum/handlers"
	"forum/models"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/Post/CreatePost", models.CreateNewPost)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/static/", handlers.StaticHandlers)
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal("could not init db:", err)
	}

	defer db.Close()
	fmt.Println("server is start in http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
