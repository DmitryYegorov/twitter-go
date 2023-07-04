package interfaces

import "twitter-go/entity"

type PostRepository interface {
	Create(data entity.CreatePostRecord) (int, error)
	FindOne(id int) (*entity.Post, error)
}
