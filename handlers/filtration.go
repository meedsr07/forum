package handlers

import (
	"net/http"

	"forum/database"
	"forum/models"
)

// Get user ID from the session token in the request cookies

func GetUserID(r *http.Request) (int, error) {
	// Get the cookie named session_token from the user's request
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}
	// take the value of the cookie
	token := cookie.Value

	var userID int
	// Query the database to find the user_id associated with the session_token
	err = database.DB.QueryRow("SELECT user_id FROM user_sessions  WHERE session_token = ?", token).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the URL path is not "/"
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusText(400), 400)
		return
	}
	// methode not post return

	//-----------------------
	userID, sessionErr := GetUserID(r)
	// get the filter query parameter from the URL
	// ------------------------- walid ---------------------------
	category := r.URL.Query().Get("category")
	if category != "" {
		FilterPostsHandler(w, r, category)
		return
	}
	// ------------------------- end walid ---------------------------
	filter := r.URL.Query().Get("filter")

	var posts []models.Post
	var err error

	switch filter {
	case "myposts":
		if sessionErr != nil {
			// If the user is not logged in, redirect them to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetMyPosts(userID)
		if err != nil {
			ErrorHandler(w, http.StatusText(500), 500)
			return
		}

	case "liked":
		if sessionErr != nil {
			// If the user is not logged in, redirect them to the login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		posts, err = database.GetLikedPosts(userID)
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}
		// default case: show all posts
	default:
		posts, err = database.Getallpost()
		if err != nil {
			ErrorHandler(w, "internal server error", 500)
			return
		}
	}

	// Create a PageData struct to pass to the template
	pageData := models.PageData{
		Posts: posts,
	}
	// If the user is logged in, IsLoggedIn == true and get the username
	if sessionErr == nil {
		var username string
		// slect the username from the users table by userID
		dbErr := database.DB.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
		// if we found the username, set IsLoggedIn to true and pass the username to the struct
		if dbErr == nil {
			pageData.IsLoggedIn = true
			pageData.Username = username
		}
	}
	///  ----------------- walid --------------------------------
	pageData.Category, err = database.GetAllCategories()
	///  ----------------- walid --------------------------------
	// Execute the template with the pageData struct
	tmpl.ExecuteTemplate(w, "index.html", pageData)
	if err != nil {
		ErrorHandler(w, http.StatusText(500), 500)
		return
	}

	tmpl.Execute(w, pageData)
}
