package main

import (
	"forum/database"
)

func main() {
	database.InitializeDB()
	err := database.SeedDB(database.DB)
	if err != nil {
		return
	}
}
