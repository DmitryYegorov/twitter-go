package interfaces

import (
	"twitter-go/entity"
	"twitter-go/utils"
)

type PostService interface {
	CreateNewPost(data entity.CreatePostRequest, userId int) (int, *utils.HttpError)
}
