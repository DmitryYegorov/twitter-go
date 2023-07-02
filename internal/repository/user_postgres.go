package repository

import (
	"github.com/jackc/pgx"
	"twitter-go/entity"
)

type UserRepoImpl struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserRepoImpl {
	return &UserRepoImpl{
		db: db,
	}
}

func (r *UserRepoImpl) Create(data entity.RegisterUserRequest) (int, error) {
	var user_id int
	err := r.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", data.Name, data.Email, data.Password).Scan(&user_id)
	if err != nil {
		return 0, err
	}

	return user_id, err
}

func (r *UserRepoImpl) FindOne(id int) (entity.User, error) {
	return entity.User{}, nil
}

func (r *UserRepoImpl) FindAll() []entity.User {
	return []entity.User{}
}

func (r *UserRepoImpl) Delete(id int) error {
	return nil
}

func (r *UserRepoImpl) Update(id int, data entity.User) (entity.User, error) {
	return entity.User{}, nil
}
