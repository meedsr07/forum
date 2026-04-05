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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("server is start in http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
