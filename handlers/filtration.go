package handlers

import (
	"fmt"
	"net/http"

	"forum/database"
)

func GetUserID(r *http.Request) (int, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return 0, err
	}
	token := cookie.Value

	var userID int
	err = database.DB.QueryRow("SELECT user_id FROM user_sessions  WHERE session_token = ?", token).Scan(&userID)
	if err != nil {
		fmt.Println("error ")
		return 0, err
	}
	return userID, nil
}

// func Feltring(w http.ResponseWriter, r *http.Request) {
// 	UserId, err := GetUserID(r)
// 	if err != nil {
// 		return
// 	}
// }
