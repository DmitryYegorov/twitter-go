package interfaces

import (
	"twitter-go/entity"
)

type PostService interface {
	CreateNewPost(data entity.CreatePostRequest, userId int) (int, error)
	GetUserPosts(userId int) ([]entity.Post, error)
}
