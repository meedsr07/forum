package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// linking go with database

var DB *sql.DB

func InitializeDB() {
	var err error
	// open a connection betwen go and database
	DB, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Read schema file
	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatal("Error reading schema.sql:", err)
	}

	// Execute schema
	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatal("Error executing schema:", err)
	}
}
