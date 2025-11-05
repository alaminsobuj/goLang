package models

import "github.com/alaminsobuj/goLang/internal/db"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() ([]User, error) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users") // ✅ now works
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func InsertUser(name, email string) error {
	_, err := db.DB.Exec("INSERT INTO users(name,email) VALUES(?,?)", name, email)
	return err
}
