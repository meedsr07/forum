package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// ─────────────────────────────────────────────
// INIT
// ─────────────────────────────────────────────

// InitDB opens (or creates) the SQLite file and runs schema.sql to create all tables.
// Call this once in main.go:
//
//	db, err := database.InitDB("forum.db")
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	// Enable foreign key enforcement (SQLite disables it by default)
	if _, err = db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	// Read and execute schema.sql
	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("read schema: %w", err)
	}
	if _, err = db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("run schema: %w", err)
	}

	DB = db
	return DB, nil
}
