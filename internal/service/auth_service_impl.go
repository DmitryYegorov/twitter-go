package service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"twitter-go/entity"
	"twitter-go/internal/repository"
	utils "twitter-go/utils"
)

type (
	AuthServiceImpl struct {
		repo *repository.Repository
	}
)

func NewAuthService(repo *repository.Repository) *AuthServiceImpl {
	return &AuthServiceImpl{repo: repo}
}

func (s *AuthServiceImpl) Login(email string, password string) (*utils.JwtResponse, *utils.HttpError) {
	foundUser, _ := s.repo.UserRepo.FindByEmail(email)

	if foundUser == nil || foundUser.EmailVerifiedAt == nil {
		return nil, utils.NewHttpError(http.StatusBadRequest, "Incorrect credentials")
	}

	logrus.Printf("found user: %+v", foundUser)

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, utils.NewHttpError(http.StatusBadRequest, "Incorrect credentials")
	}

	payload := utils.UserPayload{Id: foundUser.Id, Name: foundUser.Name, Email: foundUser.Email}
	jwtRes, err := utils.GenerateJwt(payload)
	if err != nil {
		return nil, utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return jwtRes, nil
}

func (s *AuthServiceImpl) Register(data entity.RegisterUserRequest) (int, *utils.HttpError) {
	var existingUser *entity.User = nil

	existingUser, err := s.repo.UserRepo.FindByEmail(data.Email)
	if err == nil && existingUser != nil {
		logrus.Errorf("User with email = %s already exists", existingUser.Email)
		return 0, utils.NewHttpError(http.StatusBadRequest, "user with such email already exists")
	}

	existingUser, err = s.repo.UserRepo.FindByName(data.Name)
	if err == nil && existingUser != nil {
		logrus.Errorf("User with name = %s already exists", existingUser.Name)
		return 0, utils.NewHttpError(http.StatusBadRequest, "user with such name already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)
	data.Password = string(hashedPassword)

	userId, err := s.repo.UserRepo.Create(data)

	if err != nil {
		logrus.Error(err)
		return 0, utils.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return userId, nil
}
