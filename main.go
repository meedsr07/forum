package main

import (
	"fmt"
	"forum/database"
	"forum/handlers"
	"log"
	"net/http"
)

func main() {
	database.InitializeDB()
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("server is start in http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

