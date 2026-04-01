package main

import (
	"fmt"

	"forum/database"
	"forum/handlers"
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
	likedPost , err :=  database.GetLikedPosts(database.DB , 1)
	if err != nil {
		fmt.Println("error in geting liked postes")
		return
	}
	fmt.Println(likedPost)
	usersesion , err := handlers.GetUserSession(database.DB)
	if err != nil {
		return
	}
	fmt.Println(usersesion)
}
