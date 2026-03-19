package main

import (
	"log"
	"net/http"

	"forum/database"
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
}
