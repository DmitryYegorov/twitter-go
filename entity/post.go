package entity

import "time"

type Post struct {
	Id        int        `json:"id"`
	Text      string     `json:"text"`
	CreatedBy int        `json:"createdBy"`
	CreatedAt *time.Time `json:"createdAt"`
}

type CreatePostRequest struct {
	Text string `json:"text"`
}

type CreatePostRecord struct {
	Text      string `json:"text"`
	CreatedBy int    `json:"createdBy"`
}
