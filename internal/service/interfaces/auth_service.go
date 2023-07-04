package interfaces

import (
	"twitter-go/entity"
	errors2 "twitter-go/utils"
)

type AuthService interface {
	Login(email string, password string) (interface{}, error)
	Register(data entity.RegisterUserRequest) (int, *errors2.HttpError)
}
