package interfaces

import "twitter-go/entity"

type UserRepo interface {
	Create(data entity.RegisterUserRequest) (int, error)
	FindOne(id int) (entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByName(name string) (*entity.User, error)
	FindAll() []entity.User
	Delete(id int) error
	Update(id int, data entity.User) (entity.User, error)
}
