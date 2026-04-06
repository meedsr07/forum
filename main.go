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
	
	http.HandleFunc("/static/", handlers.StaticHandlers)
	fmt.Println("server is start in http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
