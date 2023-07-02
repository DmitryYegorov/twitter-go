package repository

import (
	"github.com/jackc/pgx"
	"twitter-go/internal/repository/interfaces"
)

type Repository struct {
	UserRepo interfaces.UserRepo
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		UserRepo: NewUserPostgres(db),
	}
}
