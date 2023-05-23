package config

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func PostgresConnect(host string, port uint16, database string, user string, password string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to DB: %v\n", err)
		os.Exit(1)
	}

	return conn, err
}
