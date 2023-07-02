package service

import (
	"twitter-go/entity"
	"twitter-go/internal/repository"
)

type (
	AuthServiceImpl struct {
		repo *repository.Repository
	}
)

func NewAuthService(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Login(email string, password string) (interface{}, error) {
	return nil, nil
}

func (s *AuthServiceImpl) Register(data entity.RegisterUserRequest) (int, error) {
	return 0, nil
}
