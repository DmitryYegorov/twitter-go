package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	godotenv.Load(".env")

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, world!")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
