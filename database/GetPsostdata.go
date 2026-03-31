package database

import (
	"database/sql"

	"forum/models"
)

func Getallpost(DB *sql.DB) ([]models.Post, error) {
	var AllPost []models.Post

	rows, err := DB.Query("SELECT id, user_id, title, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// moving in rows one by one
	for rows.Next() {
		// keep going until no more rows
		var p models.Post
		err := rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Content , &p.Created_At)
		if err != nil {
			return nil, err
		}
		AllPost = append(AllPost, p)

	}
	return AllPost, nil
}


func GetMyPosts(DB *sql.DB, userID int) ([]models.Post, error) {
	var UserPosts []models.Post

	rows, err := DB.Query("SELECT id, user_id, title, content, created_at FROM posts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.Id, &p.UserID, &p.Title, &p.Content, &p.Created_At)
		if err != nil {
			return nil, err
		}
		UserPosts = append(UserPosts, p)
	}

	return UserPosts, nil
}