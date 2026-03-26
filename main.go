package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// 1. Serve static files (CSS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Serve the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Parse  the HTML file
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute (Send) the HTML to the user's browser
		tmpl.Execute(w, nil)
	})

	// 3. Start the server on port 8080
	fmt.Println(" Server is running! Open your browser and go to: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}
