package handlers

import (
	"bytes"
	"html/template"
	"net/http"
)

// ErrorPage holds the data to be displayed on the error page
type ErrorPage struct {
	Code    int
	Message string
}

// ErrorHandler renders a custom error page with the given message and status code
func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {

	// Create the data for the error page
	errorPage := ErrorPage{
		Code:    statusCode,
		Message: message,
	}

	// Parse the error template file
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// If the template cannot be parsed, return a generic 500 error to avoid infinite recursion
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template into a buffer first to check for errors
	var buff bytes.Buffer
	if err := tmpl.Execute(&buff, errorPage); err != nil {
		// If the template cannot be executed, return a generic 500 error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the HTTP status code and the buffered content to the response
	w.Write(buff.Bytes())
}
