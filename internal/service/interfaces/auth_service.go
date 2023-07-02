package interfaces

import "twitter-go/entity"

type AuthService interface {
	Login(email string, password string) (interface{}, error)
	Register(data entity.RegisterUserRequest) (int, error)
}
