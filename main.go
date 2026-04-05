package main

import (
	"fmt"
	"forum/database"
	"forum/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/Post/CreatePost", handlers.CreateNewPost)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/static/", handlers.StaticHandlers)
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal("could not init db:", err)
	}

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	fmt.Printf("server is start in http://localhost:%s\n", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
