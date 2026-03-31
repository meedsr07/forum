package main

import (
	"fmt"

	"forum/database"
)

func main() {
	database.InitializeDB()
	Allpost, err := database.Getallpost(database.DB)
	if err != nil {
		fmt.Println("error in reading database")
		return
	}
	fmt.Println(Allpost)
	UserPosts , err := database.GetMyPosts(database.DB , 1)
	if err  !=  nil {
		fmt.Println("error in git user posts")
		return
	}
	fmt.Println(UserPosts)
}
