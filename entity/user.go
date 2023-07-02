package entity

import "time"

type User struct {
	id                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	email_verified_at time.Time `json:"email_Verified_At"`
	created_at        time.Time `json:"created_At"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
