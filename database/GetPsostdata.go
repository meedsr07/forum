package database

import (
	"forum/models"
)

func Getallpost() ([]models.Post, error) {
	var AllPost []models.Post
	// asking the database to get data from the posts table
	rows, err := DB.Query(`
    SELECT posts.id, posts.user_id, users.username, posts.title, posts.content, posts.created_at 
    FROM posts 
    JOIN users ON posts.user_id = users.id
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// moving in rows one by one
	for rows.Next() {
		// keep going until no more rows
		var p models.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Username, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		AllPost = append(AllPost, p)

	}
	return AllPost, nil
}

func GetMyPosts(userID int) ([]models.Post, error) {
	var UserPosts []models.Post

	rows, err := DB.Query("SELECT id, user_id, Username , title, content, created_at FROM posts WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		UserPosts = append(UserPosts, p)
	}

	return UserPosts, nil
}

func GetOnePost(postID int) (models.Post, error) {
	var post models.Post

	err := DB.QueryRow("SELECT id, user_id, title, content, created_at FROM posts WHERE id = ?",postID,).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
	)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetLikedPosts(userID int) ([]models.Post, error) {
	var LikedPosts []models.Post

	rows, err := DB.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.content, posts.created_at
		FROM posts
		JOIN post_reactions ON posts.id = post_reactions.post_id
		WHERE post_reactions.user_id = ?
		AND post_reactions.reaction = 1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Username, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		LikedPosts = append(LikedPosts, p)
	}

	return LikedPosts, nil
}
