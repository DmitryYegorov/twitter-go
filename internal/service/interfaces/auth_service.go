package interfaces

import (
	"twitter-go/entity"
	"twitter-go/utils"
)

type AuthService interface {
	Login(email string, password string) (*utils.JwtResponse, *utils.HttpError)
	Register(data entity.RegisterUserRequest) (int, *utils.HttpError)
}
