package database

import (
	"database/sql"
	"fmt"

	"forum/models"
)

func Getallpost() ([]models.Post, error) {
	var AllPost []models.Post
	// asking the database to get data from the posts table and username
	rows, err := DB.Query(`
    SELECT posts.id, posts.user_id, users.username, posts.title, posts.content, posts.created_at 
    FROM posts 
    JOIN users ON posts.user_id = users.id
	ORDER BY posts.created_at DESC
	
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
	// Select all posts from the database that belong to a specific userID
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
		err := rows.Scan(&post.ID,&post.UserID,&post.Title,&post.Content,&post.CreatedAt,&post.Username)
		
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
		err := rows.Scan(&post.ID,&post.UserID,&post.Title,&post.Content,&post.CreatedAt,&post.Username)
	
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// ------------------------------------------------ walid

// GetPostVotes returns the like and dislike count for a given post.
func GetPostVotes(postID int) (int, int) {
	var likes, dislikes int

	DB.QueryRow(`
		SELECT 
			COALESCE(SUM(CASE WHEN value = 1 THEN 1 END),0),
			COALESCE(SUM(CASE WHEN value = -1 THEN 1 END),0)
		FROM likes
		WHERE post_id=?`,
		postID,
	).Scan(&likes, &dislikes)

	return likes, dislikes
}

// GetPostsByCategory returns all posts that belong to a given category ID.
func GetPostsByCategory(categoryID int) ([]models.Post, error) {
	var posts []models.Post
	var exists int

	err := DB.QueryRow("SELECT 1 FROM categories WHERE id = ?", categoryID).Scan(&exists)

	if err == sql.ErrNoRows {
		return nil , fmt.Errorf("category not found")
	}
	if err != nil {
		return nil , err
	}
	rows, err := DB.Query(`
		SELECT posts.id, posts.user_id, users.username, posts.title, posts.content, posts.created_at
		FROM posts
		JOIN post_categories ON posts.id = post_categories.post_id
		JOIN users ON posts.user_id = users.id
		WHERE post_categories.category_id = ?
		ORDER BY posts.created_at DESC
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Username,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("faild to scan post : %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetAllCategories returns all categories from the database.
func GetAllCategories() ([]models.Category, error) {
	rows, err := DB.Query(`SELECT id, name FROM categories ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, rows.Err()
}

// CategoryExists checks whether a category with the given ID exists.
func CategoryExists(categoryID int) (bool, error) {
	var id int
	err := DB.QueryRow(`SELECT id FROM categories WHERE id = ?`, categoryID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetCommentVotes(commentID int) (int, int) {
	var likes, dislikes int

	DB.QueryRow(`
		SELECT 
			COALESCE(SUM(CASE WHEN value = 1 THEN 1 END),0),
			COALESCE(SUM(CASE WHEN value = -1 THEN 1 END),0)
		FROM likes
		WHERE comment_id=?`,
		commentID,
	).Scan(&likes, &dislikes)

	return likes, dislikes
}
