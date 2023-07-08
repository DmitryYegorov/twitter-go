package handlers

import (
	"github.com/caarlos0/env/v9"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"twitter-go/config"
	"twitter-go/internal/handlers/middlewares"
	"twitter-go/internal/service"
)

type Handler struct {
	AuthHandler
	PostsHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		AuthHandler:  &authHandler{AuthService: service.AuthService},
		PostsHandler: &postsHandler{PostsService: service.PostService},
	}
}

func SetApi(e *echo.Echo, h *Handler) {

	jwtConfig := config.JwtAuthConfig{}
	if err := env.Parse(&jwtConfig); err != nil {
		logrus.Fatalf("JwtConfig is not loaded")
	}

	group := e.Group("/api")

	authGroup := group.Group("/auth")
	{
		authGroup.POST("/login/email", h.AuthHandler.Login)
		authGroup.POST("/register", h.AuthHandler.Register)
	}

	postGroup := group.Group("/posts")
	{
		postGroup.POST("", h.PostsHandler.Create, middlewares.JwtAuth(jwtConfig.AccessSecretKey))
		postGroup.GET("/user/:userId", h.PostsHandler.GetAllByUser)
	}
}

func Echo() *echo.Echo {
	e := echo.New()

	return e
}
