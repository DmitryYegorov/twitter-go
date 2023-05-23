package entities

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	isActive  bool      `json:"isActive"`
	createdAt time.Time `json:"createdAt"`
}
