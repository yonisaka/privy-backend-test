package domain

import "time"

type User struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username" binding:"required"`
	Email     string     `json:"email"`
	Password  string     `json:"password" binding:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Auth struct {
	Access_token string `json:"access_token"`
	Expired_at   string `json:"expired_at"`
}
