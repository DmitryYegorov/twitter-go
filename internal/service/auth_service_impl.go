package service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"twitter-go/entity"
	"twitter-go/internal/repository"
	errors2 "twitter-go/utils"
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

func (s *AuthServiceImpl) Register(data entity.RegisterUserRequest) (int, *errors2.HttpError) {
	var existingUser *entity.User = nil

	existingUser, err := s.repo.UserRepo.FindByEmail(data.Email)
	if err == nil && existingUser != nil {
		logrus.Errorf("User with email = %s already exists", existingUser.Email)
		return 0, errors2.NewHttpError(http.StatusBadRequest, "user with such email already exists")
	}

	existingUser, err = s.repo.UserRepo.FindByName(data.Name)
	if err == nil && existingUser != nil {
		logrus.Errorf("User with name = %s already exists", existingUser.Name)
		return 0, errors2.NewHttpError(http.StatusBadRequest, "user with such name already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(hashedPassword)

	userId, err := s.repo.UserRepo.Create(data)

	if err != nil {
		logrus.Error(err)
		return 0, errors2.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return userId, nil
}
