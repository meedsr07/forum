package handlers

import (
	"database/sql"
	"forum/models"
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