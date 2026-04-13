package database

import (
	"forum/models"
)

func Getallpost() ([]models.Post, error) {
	var AllPost []models.Post
	rows, err := DB.Query(`
		SELECT p.id, p.user_id, u.username, p.title, p.content, p.created_at,
		       COALESCE(c.name, '')
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN categories c ON p.category_id = c.id
		ORDER BY p.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Username, &p.Title, &p.Content, &p.CreatedAt, &p.Category)
		if err != nil {
			return nil, err
		}
		p.Likes, p.Dislikes = GetPostVotes(p.ID)
		AllPost = append(AllPost, p)
	}
	return AllPost, rows.Err()
}

func GetMyPosts(userID int) ([]models.Post, error) {
	var posts []models.Post

	rows, err := DB.Query(`
		SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.username,
		       COALESCE(c.name, '')
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.user_id = ?
		ORDER BY p.created_at DESC
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
			&post.Category,
		)
		if err != nil {
			return nil, err
		}
		post.Likes, post.Dislikes = GetPostVotes(post.ID)
		posts = append(posts, post)
	}

	return posts, rows.Err()
}

func GetOnePost(postID int) (models.Post, error) {
	var post models.Post

	err := DB.QueryRow(`
		SELECT p.id, p.user_id, u.username, p.title, p.content, p.created_at,
		       COALESCE(c.name, '')
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = ?
	`, postID).Scan(
		&post.ID,
		&post.UserID,
		&post.Username,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.Category,
	)
	if err != nil {
		return models.Post{}, err
	}

	post.Likes, post.Dislikes = GetPostVotes(post.ID)
	return post, nil
}

func GetLikedPosts(userID int) ([]models.Post, error) {
	var posts []models.Post

	rows, err := DB.Query(`
		SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.username,
		       COALESCE(c.name, '')
		FROM posts p
		JOIN likes l ON p.id = l.post_id
		JOIN users u ON p.user_id = u.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE l.user_id = ? AND l.value = 1
		ORDER BY p.created_at DESC
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
			&post.Category,
		)
		if err != nil {
			return nil, err
		}
		post.Likes, post.Dislikes = GetPostVotes(post.ID)
		posts = append(posts, post)
	}

	return posts, rows.Err()
}

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

	rows, err := DB.Query(`
		SELECT p.id, p.user_id, u.username, p.title, p.content, p.created_at,
		       COALESCE(c.name, '')
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.category_id = ?
		ORDER BY p.created_at DESC
	`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Username, &post.Title, &post.Content, &post.CreatedAt, &post.Category)
		if err != nil {
			return nil, err
		}
		post.Likes, post.Dislikes = GetPostVotes(post.ID)
		posts = append(posts, post)
	}

	return posts, rows.Err()
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
