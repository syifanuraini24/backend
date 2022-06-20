package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
func GetUser(count int) ([]User, error) {

	rows, err := DB.Query("SELECT id, first_name, last_name, email, ip_address from people LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	user := make([]User, 0)

	for rows.Next() {
		singleUser := User{}
		err = rows.Scan(&singleUser.Id, &singleUser.FirstName, &singleUser.LastName, &singleUser.Email)

		if err != nil {
			return nil, err
		}

		user = append(user, singleUser)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return user, err
}
