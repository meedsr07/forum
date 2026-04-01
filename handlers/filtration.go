package handlers

import "database/sql"

type sessions struct {
	Token string
	UserID int
	Created_At string
}

func GetUserSession(DB *sql.DB) ([]sessions , error) {
	var Usersesion []sessions
	rows , err := DB.Query("SELECT token , user_id , created_at FROM  sessions ")
	if err != nil {
		return nil , err 
	}
	for rows.Next() {
		var p sessions
		err := rows.Scan(&p.Token , &p.UserID , &p.Created_At)
		if err != nil {
			return nil , err
		}
		Usersesion = append(Usersesion, p)
	}
	return  Usersesion , nil 
}