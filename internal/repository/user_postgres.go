package repository

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
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

func (r *UserRepoImpl) FindByEmail(email string) (*entity.User, error) {
	query := "SELECT u.id, u.name, u.email, u.created_at, u.email_verified_at, u.password FROM users AS u WHERE u.email = $1"
	row := r.db.QueryRow(query, email)

	logrus.Printf("input email %s", email)

	user := &entity.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.EmailVerifiedAt, &user.Password)

	logrus.Printf("%+v", err)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Пользователь не найден
		}
		return nil, err // Возникла ошибка при выполнении запроса
	}

	if user.Id == 0 {
		return nil, nil
	}

	return user, nil // Пользователь найден
}

func (r *UserRepoImpl) FindByName(name string) (*entity.User, error) {
	query := "SELECT u.id, u.name, u.email, u.created_at, u.email_verified_at FROM users AS u WHERE u.name = $1"
	row := r.db.QueryRow(query, name)

	user := &entity.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.EmailVerifiedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Пользователь не найден
		}
		return nil, err // Возникла ошибка при выполнении запроса
	}

	return user, nil // Пользователь найден
}
