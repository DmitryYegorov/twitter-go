package entities

import "time"

type Post struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int       `json:"createdBy"`
}
