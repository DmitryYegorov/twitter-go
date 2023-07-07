package service

import (
	"net/http"
	"twitter-go/entity"
	"twitter-go/internal/repository"
	"twitter-go/utils"
)

type PostServiceImpl struct {
	repo *repository.Repository
}

func NewPostsService(repo *repository.Repository) *PostServiceImpl {
	return &PostServiceImpl{repo: repo}
}

func (s *PostServiceImpl) CreateNewPost(data entity.CreatePostRequest, userId int) (int, *utils.HttpError) {
	postId, err := s.repo.PostRepo.Create(entity.CreatePostRecord{
		Text:      data.Text,
		CreatedBy: userId,
	})

	if err != nil {
		return 0, utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	return postId, nil
}
