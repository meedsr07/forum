package database

import "database/sql"

func SeedDB(db *sql.DB) error {
	// ── Categories ─────────────────────────────────────────────────────────
	categories := []string{"General", "Tech", "Gaming", "Movies", "Science"}
	for _, c := range categories {
		_, err := db.Exec(`INSERT OR IGNORE INTO categories (name) VALUES (?)`, c)
		if err != nil {
			return err
		}
	}
	return nil
}
