package main

import (
	"database/sql"
	"fmt"

	"forum/database"
)

func main() {
	DB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return
	}
	err1 := database.SeedDB(DB)
	if err1 != nil {
		fmt.Println("a")
		return
	}
}
