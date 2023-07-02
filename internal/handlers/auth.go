package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"twitter-go/entity"
	"twitter-go/internal/service/interfaces"
)

type (
	AuthHandler interface {
		Login(c echo.Context) error
		Register(c echo.Context) error
	}

	authHandler struct {
		AuthService interfaces.AuthService
	}
)

func (h *authHandler) Login(c echo.Context) error {
	//h.Service.AuthService.Login()
	return nil
}

func (h *authHandler) Register(c echo.Context) error {
	var req entity.RegisterUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Register error")
	}

	user_id, err := h.AuthService.Register(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error")
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"user_id": user_id,
	})
}

//
//func RegisterUser(c echo.Context) error {
//	var req *entity.RegisterUserRequest
//
//	err := c.Bind(&req)
//
//	return err
//}
//
//func LoginUser(c echo.Context) error {
//	return nil
//}
