package entities

import "time"

type Comment struct {
	Id        int       `json:"id"`
	Content   string    `json:"content"`
	PostId    int       `json:"postId"`
	CreatedBy int       `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
}
