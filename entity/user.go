package entity

import "time"

type User struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	EmailVerifiedAt *time.Time `json:"email_Verified_At"`
	CreatedAt       time.Time  `json:"created_At"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
