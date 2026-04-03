package database

import "database/sql"

// SeedDB populates the database with realistic test data.
// Call this once after InitDB during development:
//
//	database.SeedDB(db)
//
// Safe to call multiple times — uses INSERT OR IGNORE so it won't duplicate data.
func SeedDB(db *sql.DB) error {
	// ── Users ──────────────────────────────────────────────────────────────
	// Passwords are plain text here for dev convenience.
	// In production Person 1 will hash them with bcrypt.
	users := []struct{ email, username, password string }{
		{"alice@example.com", "alice", "password123"},
		{"bob@example.com", "bob", "password123"},
		{"charlie@example.com", "charlie", "password123"},
		{"diana@example.com", "diana", "password123"},
	}
	for _, u := range users {
		_, err := db.Exec(
			`INSERT OR IGNORE INTO users (email, username, password) VALUES (?, ?, ?)`,
			u.email, u.username, u.password,
		)
		if err != nil {
			return err
		}
	}

	// ── Categories ─────────────────────────────────────────────────────────
	categories := []string{"General", "Tech", "Gaming", "Movies", "Science"}
	for _, c := range categories {
		_, err := db.Exec(`INSERT OR IGNORE INTO categories (name) VALUES (?)`, c)
		if err != nil {
			return err
		}
	}

	// ── Posts ──────────────────────────────────────────────────────────────
	posts := []struct {
		userID  int
		title   string
		content string
	}{
		{1, "Welcome to the Forum!", "Hey everyone, glad to have you here. Feel free to post anything."},
		{2, "Best programming languages in 2024", "I think Go is underrated. What do you all think?"},
		{1, "Anyone else watching the new season?", "Just finished episode 3, no spoilers but wow."},
		{3, "Cool science fact of the day", "Did you know honey never expires? Found 3000 year old honey in Egypt."},
		{4, "Recommend me a game", "Looking for something chill to play on weekends. Any suggestions?"},
		{2, "Go vs Rust – which one to learn?", "I've been going back and forth. Both seem awesome for systems programming."},
	}
	for _, p := range posts {
		_, err := db.Exec(
			`INSERT OR IGNORE INTO posts (user_id, title, content) VALUES (?, ?, ?)`,
			p.userID, p.title, p.content,
		)
		if err != nil {
			return err
		}
	}

	// ── Post ↔ Category links ──────────────────────────────────────────────
	// post_id → category_id  (using the order we inserted above)
	postCats := [][2]int{
		{1, 1}, // Welcome → General
		{2, 2}, // Programming → Tech
		{3, 4}, // Season → Movies
		{4, 5}, // Science fact → Science
		{5, 3}, // Game rec → Gaming
		{6, 2}, // Go vs Rust → Tech
		{5, 1}, // Game rec also in General
	}
	for _, pc := range postCats {
		db.Exec(
			`INSERT OR IGNORE INTO post_categories (post_id, category_id) VALUES (?, ?)`,
			pc[0], pc[1],
		)
	}

	// ── Comments ───────────────────────────────────────────────────────────
	// This is YOUR section! These are the comments you will be querying.
	comments := []struct {
		postID  int
		userID  int
		content string
	}{
		{1, 2, "Thanks for the warm welcome, happy to be here!"},
		{1, 3, "Great forum, looking forward to some good discussions."},
		{2, 1, "Go is amazing for backend stuff, especially concurrency."},
		{2, 3, "Rust is harder to learn but the performance is insane."},
		{2, 4, "I started with Python, but Go feels much more production-ready."},
		{3, 2, "No spoilers please! I'm only on episode 1 haha"},
		{3, 4, "Episode 3 was wild, can't wait for the next one."},
		{4, 1, "That's so cool! What about coffee – does that expire?"},
		{4, 2, "Bees are basically nature's preservatives."},
		{5, 1, "Try Stardew Valley, super chill and relaxing."},
		{5, 3, "Minecraft creative mode is my go-to for relaxing."},
		{6, 4, "Go for simplicity, Rust for safety. Depends on the project."},
	}
	for _, c := range comments {
		_, err := db.Exec(
			`INSERT OR IGNORE INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`,
			c.postID, c.userID, c.content,
		)
		if err != nil {
			return err
		}
	}

	// ── Likes / Dislikes ───────────────────────────────────────────────────
	// value: 1 = like, -1 = dislike
	likes := []struct {
		userID, postID, commentID, value int
	}{
		// Post likes
		{2, 1, 0, 1},
		{3, 1, 0, 1},
		{4, 1, 0, 1},
		{1, 2, 0, 1},
		{3, 2, 0, 1},
		{4, 2, 0, -1},
		{1, 4, 0, 1},
		{2, 4, 0, 1},
		{3, 5, 0, 1},
		{1, 6, 0, 1},
		{4, 6, 0, 1},
		// Comment likes
		{1, 0, 1, 1},
		{3, 0, 1, 1},
		{2, 0, 3, 1},
		{4, 0, 4, 1},
		{1, 0, 10, 1},
		{2, 0, 10, 1},
	}
	for _, l := range likes {
		var postIDPtr, commentIDPtr interface{}
		if l.postID != 0 {
			postIDPtr = l.postID
		}
		if l.commentID != 0 {
			commentIDPtr = l.commentID
		}
		db.Exec(
			`INSERT OR IGNORE INTO likes (user_id, post_id, comment_id, value) VALUES (?, ?, ?, ?)`,
			l.userID, postIDPtr, commentIDPtr, l.value,
		)
	}

	return nil
}

