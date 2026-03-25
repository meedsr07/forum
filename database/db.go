package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// linking go with database

var DB *sql.DB

func InitializeDB() {
	var err error
	// open a connection betwen go and database
	DB, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
}
