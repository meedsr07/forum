package handlers

import (
	"net/http"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	//cookie, err := r.Cookie("session_id")
	//if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		//return
	//}

	//var userID int
	//err = db.QueryRow("SELECT user_id FROM sessions WHERE session_id = ?", cookie.Value).Scan(&userID)
	//if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	//}
}
