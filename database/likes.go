package database

import "database/sql"

// HandleVote implements the full like/dislike toggle logic for a post.
//   - If no existing vote → INSERT
//   - If same vote exists → DELETE (toggle off)
//   - If opposite vote exists → UPDATE
func HandleVote(userID, postID, value int) error {
	var existingValue int
	err := DB.QueryRow(
		`SELECT value FROM likes WHERE user_id = ? AND post_id = ?`,
		userID, postID,
	).Scan(&existingValue)

	if err == sql.ErrNoRows {
		// No previous vote — insert new one
		_, err = DB.Exec(
			`INSERT INTO likes (user_id, post_id, value) VALUES (?, ?, ?)`,
			userID, postID, value,
		)
		return err
	}
	if err != nil {
		return err
	}

	if existingValue == value {
		// Same vote — toggle off (delete)
		_, err = DB.Exec(
			`DELETE FROM likes WHERE user_id = ? AND post_id = ?`,
			userID, postID,
		)
		return err
	}

	// Opposite vote — switch
	_, err = DB.Exec(
		`UPDATE likes SET value = ? WHERE user_id = ? AND post_id = ?`,
		value, userID, postID,
	)
	return err
}

func HandleVotecomment(userID, commentID, value int) error {
	var existingValue int
	err := DB.QueryRow(
		`SELECT value FROM likes WHERE user_id = ? AND comment_id= ?`,
		userID, commentID,
	).Scan(&existingValue)

	if err == sql.ErrNoRows {
		// No previous vote — insert new one
		_, err = DB.Exec(
			`INSERT INTO likes (user_id, comment_id, value) VALUES (?, ?, ?)`,
			userID, commentID, value,
		)
		return err
	}
	if err != nil {
		return err
	}

	if existingValue == value {
		// Same vote — toggle off (delete)
		_, err = DB.Exec(
			`DELETE FROM likes WHERE user_id = ? AND comment_id = ?`,
			userID, commentID,
		)
		return err
	}

	// Opposite vote — switch
	_, err = DB.Exec(
		`UPDATE likes SET value = ? WHERE user_id = ? AND comment_id = ?`,
		value, userID, commentID,
	)
	return err
}
