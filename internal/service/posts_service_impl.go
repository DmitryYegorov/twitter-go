package service

import (
	"github.com/sirupsen/logrus"
	"twitter-go/entity"
	"twitter-go/internal/repository"
)

type PostServiceImpl struct {
	repo *repository.Repository
}

func NewPostsService(repo *repository.Repository) *PostServiceImpl {
	return &PostServiceImpl{repo: repo}
}

func (s *PostServiceImpl) CreateNewPost(data entity.CreatePostRequest, userId int) (int, error) {
	postId, err := s.repo.PostRepo.Create(entity.CreatePostRecord{
		Text:      data.Text,
		CreatedBy: userId,
	})

	if err != nil {
		return 0, err
	}
	return postId, nil
}

func (s *PostServiceImpl) GetUserPosts(userId int) ([]entity.Post, error) {
	list, err := s.repo.PostRepo.FindByUserId(userId)
	if err != nil {
		logrus.Fatalf(err.Error())
		return nil, err
	}

	return list, nil
}
