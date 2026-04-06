package handlers

import (
	"net/http"
	"os"
)

func StaticHandlers(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[1:]

	// Verify if the requested file exists on the server.
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		// If the file is missing or a directory is requested, return a custom 404 error page.
		ErrorHandler(w, "Resource Not Found", http.StatusNotFound)
		return
	}

	// Serve the static file using the standard FileServer after stripping the "/static/" prefix.
	staticServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	staticServer.ServeHTTP(w, r)
}
