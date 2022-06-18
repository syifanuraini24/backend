package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

type ProfileErrorResponse struct {
	Error string `json:"error"`
}

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileSuccesResponse struct {
	Profile []Profile `json:"profile"`
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

func (u *UserRepository) Register(nama string, email string, password string) error {
	_, err := u.db.Exec("INSERT INTO users (nama, email, password) VALUES (?, ?, ?)", nama, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateUser(id string, nama string, email string, password string, role string, loggedin bool) error {
	_, err := u.db.Exec("UPDATE users SET nama = ?, email = ?, password = ?, role = ?, loggedin = ? WHERE id = ?", nama, email, password, role, loggedin, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) DeleteUser(id string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) CheckUser(email string) (string, error) {
	var user User
	sqlStatement := "SELECT email FROM users WHERE email = ?"

	err := u.db.QueryRow(sqlStatement, email).Scan(&user.Email)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (p *UserRepository) GetProfile() ([]User, error) {
	var profiles []User

	rows, err := p.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var profile User
		err := rows.Scan(&profile.ID, &profile.Nama, &profile.Email, &profile.Password, &profile.Role, &profile.Loggedin)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
