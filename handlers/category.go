package handlers

import (
	"database/sql"
	"forum/database"
	"forum/models"
	"html/template"
	"net/http"
	"strconv"
)

func GetAllCategories(db *sql.DB) ([]models.Category, error) {
	query := `SELECT id, name FROM categories ORDER BY id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func SeedCategories(db *sql.DB) error {
	categories := []string{"General", "Tech", "Sport", "News"}

	for _, name := range categories {
		_, err := db.Exec(`
            INSERT INTO categories (name)
            SELECT ?
            WHERE NOT EXISTS (
                SELECT 1 FROM categories WHERE name = ?
            )
        `, name, name)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetPostsByCategory(db *sql.DB, categoryID int) ([]models.Post, error) {
	query := `
        SELECT p.id, p.title, p.content, c.name
        FROM posts p
        JOIN categories c ON p.category_id = c.id
        WHERE p.category_id = ?
        ORDER BY p.id DESC
    `

	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var p models.Post

		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Category)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func FilterPostsHandler(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := r.URL.Query().Get("category_id")

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		ErrorHandler(w, "Invalid category", 400)
		return
	}

	posts, err := GetPostsByCategory(database.DB, categoryID)
	if err != nil {
		ErrorHandler(w, "Error loading posts", 500)
		return
	}

	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, posts)
}
