package repository

import "time"

type User struct {
	ID        string    `json:"user_id"`
	Nama      string    `json:"nama" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role"`
	Loggedin  bool      `json:"loggedin"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
