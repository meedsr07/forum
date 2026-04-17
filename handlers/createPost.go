package handlers

import (
	
	"net/http"
	"strconv"
	"strings"

	"forum/database"
)

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusText(404), 404)
		return
	}
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var userID int
	err = database.DB.QueryRow("SELECT user_id FROM user_sessions WHERE session_token = ?", cookie.Value).Scan(&userID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}
	RowsCategorys, err := database.DB.Query("SELECT (id) FROM categories")
	if err != nil {
		ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer RowsCategorys.Close()
	MapCategorir := make(map[string]string)
	slicIDCategores := []string{}
	for RowsCategorys.Next() {
		var id int
		err = RowsCategorys.Scan(&id)
		if err != nil {
			ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		slicIDCategores = append(slicIDCategores, strconv.Itoa(id))
		MapCategorir[strconv.Itoa(id)] = ""
	}
	err = r.ParseForm()
	if err != nil {
		ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")
	lenghtCategores := 0
	for _, v := range slicIDCategores {
		st := r.FormValue(v)
		if st != "" {
			lenghtCategores++
			MapCategorir[v] = st
		}
	}
	category := r.FormValue("category")
	title = strings.TrimSpace(title)
	content = strings.TrimSpace(content)
	category = strings.TrimSpace(category)
	if title == "" || content == "" || lenghtCategores == 0 || len(title) > 200 || len(content) > 4096 {
		ErrorHandler(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	PosT, err := database.DB.Exec(
		`INSERT INTO posts (user_id, title, content)
	 VALUES (?, ?, ?)`,
		userID, title, content,
	)
	if err != nil {
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	PostID, err := PosT.LastInsertId()
	if err != nil {
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for k, v := range MapCategorir {
		if v != "" {
			categoryID, err := strconv.Atoi(k)
			if err != nil {
				ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			_, err = database.DB.Exec(`INSERT INTO post_categories (post_id ,category_id) VALUES (?, ?)`, PostID, categoryID)
			if err != nil {
				ErrorHandler(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
