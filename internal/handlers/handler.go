package handlers

import (
	"github.com/labstack/echo/v4"
	"twitter-go/internal/service"
)

type Handler struct {
	AuthHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		AuthHandler: &authHandler{AuthService: service.AuthService},
	}
}

func SetApi(e *echo.Echo, h *Handler) {
	group := e.Group("/api")

	authGroup := group.Group("/auth")
	{
		authGroup.POST("/login/email", h.AuthHandler.Login)
		authGroup.POST("/register", h.AuthHandler.Register)
	}
}

func Echo() *echo.Echo {
	e := echo.New()

	return e
}
