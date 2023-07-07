package main

import (
	"github.com/caarlos0/env/v9"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"twitter-go/config"
	"twitter-go/internal/handlers"
	"twitter-go/internal/repository"
	"twitter-go/internal/service"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(`Failed loading of environment variables: %+v`, err)
	}

	dbConfig := &config.DataBaseConfig{}
	if err := env.Parse(dbConfig); err != nil {
		log.Fatalf("DB is not load successfully: %+v, %+v", err, dbConfig)
	}

	log.Printf("Database configuration: %+v", dbConfig)

	e := handlers.Echo()

	log.Println("Echo handlers run")

	pgxInstance, err := pgx.Connect(pgx.ConnConfig{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		User:     dbConfig.User,
		Password: dbConfig.Password,
		Database: dbConfig.Name,
	})
	defer func(pgxInstance *pgx.Conn) {
		err := pgxInstance.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(pgxInstance)
	if err != nil {
		log.Fatalf("PG start failed: %+v", err)
	}

	repo := repository.New(pgxInstance)
	services := &service.Service{
		AuthService: service.NewAuthService(repo),
		PostService: service.NewPostsService(repo),
	}
	hands := handlers.New(services)

	handlers.SetApi(e, hands)

	e.Logger.Fatal(e.Start(":4444"))
}
