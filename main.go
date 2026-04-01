package main

import (
	"fmt"
	"forum/database"
	"log"
	"net/http"
)

func main() {
	database.InitializeDB()
	http.HandleFunc("/", HomePage)
	fmt.Println("server is start in http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methode not allwed ", 405)
	}
	http.ServeFile(w, r, "templates/index.html")
}
