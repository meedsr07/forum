package handlers

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "methode not allwed ", 405)
	}
	http.ServeFile(w, r, "templates/index.html")
}
