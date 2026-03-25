package database


type Post struct {
	ID      int
	UserID  int
	Title   string
	Content string
}


func GetAllPosts() ([]Post, error) {
    rows, err := DB.Query("SELECT id, user_id, title, content FROM posts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var p Post
        if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content); err != nil {
            return nil, err
        }
        posts = append(posts, p)
    }

    return posts, nil
}