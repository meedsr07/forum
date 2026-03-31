package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/database"
	"forum/handlers"
)

func main() {
	// 1. Open the database using your teammate's InitDB function
	db, err := database.InitDB("forum.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	// Close the database safely when the server stops
	defer db.Close()

	// 2. Serve static files (CSS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 3. Serve the Login HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prevent the "/" route from stealing other URLs (like /register)
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		handlers.LoginHandler(w, r, db)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, db)
	})

	// 4. The Register Route: We send the 'db' variable to the handler
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, db)
	})
	http.HandleFunc("/home_page", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(w, r)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogoutHandler(w, r, db)
	})
	// 5. Start the server
	fmt.Println("🚀 Server is running! Open your browser and go to: http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
