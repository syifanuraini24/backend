package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var users []User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Nama, &user.Email, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
	// TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.FetchUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == username && user.Password == password {
			return &user.Email, nil
		}
	}

	return nil, fmt.Errorf("Login Failed")
	// TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	var role string

	rows, err := u.db.Query("SELECT role FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
	}

	return &role, nil
	// TODO: replace this
}
