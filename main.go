package main

import (
	"github.com/jackc/pgx"
	"os"
	"strconv"
	"twitter-go/internal/handlers"
	"twitter-go/internal/repository"
	"twitter-go/internal/service"
)

func main() {
	e := handlers.Echo()

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		//...catch
	}

	pgxInstance, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	repo := repository.New(pgxInstance)
	services := &service.Service{
		AuthService: service.NewAuthService(repo),
	}
	hands := handlers.New(services)

	handlers.SetApi(e, hands)

	e.Logger.Fatal(e.Start(":4444"))
}
