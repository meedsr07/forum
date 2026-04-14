package database

import (
	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

// CreateComment inserts a new comment into the database.
// postID  → which post this comment belongs to
// userID  → who wrote it
// content → the comment text
func CreateComment(postID, userID int, content string) error {
	_, err := DB.Exec(
		`INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`,
		postID, userID, content,
	)
	return err
}

// GetCommentsByPost returns all comments for a given post, oldest first,
// with the author's username joined from the users table.
func GetCommentsByPost(postID int) ([]models.Comment, error) {
	rows, err := DB.Query(`
			SELECT c.id, c.post_id, c.user_id, u.username, c.content, c.created_at
			FROM comments c
			JOIN users u ON c.user_id = u.id
			WHERE c.post_id = ?
			ORDER BY c.created_at ASC
		`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Username, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

