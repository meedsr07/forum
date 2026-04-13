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
	var posts []models.Post

	rows, err := DB.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.content, posts.created_at, users.username
		FROM posts
		JOIN users ON posts.user_id = users.id
		WHERE posts.user_id = ?
		ORDER BY posts.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Username, 
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetOnePost(postID int) (models.Post, error) {
	var post models.Post

	err := DB.QueryRow("SELECT id, user_id, title, content, created_at FROM posts WHERE id = ?", postID).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
	)
	if err != nil {
		return models.Post{}, err
	}
	err = DB.QueryRow("SELECT username FROM users WHERE id = ?", post.UserID).Scan(&post.Username)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func GetLikedPosts(userID int) ([]models.Post, error) {
	var posts []models.Post

	rows, err := DB.Query(`
		SELECT posts.id, posts.user_id, posts.title, posts.content, posts.created_at, users.username
		FROM posts
		JOIN likes ON posts.id = likes.post_id
		JOIN users ON posts.user_id = users.id
		WHERE likes.user_id = ?
		ORDER BY posts.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.Username, 
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
