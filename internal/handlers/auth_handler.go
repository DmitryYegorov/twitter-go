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
	var req entity.LoginEmailUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//if err := c.Validate(req); err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}

	res, err := h.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"access":  res.Access,
		"refresh": res.Refresh,
	})
}

func (h *authHandler) Register(c echo.Context) error {
	var req entity.RegisterUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userId, err := h.AuthService.Register(req)

	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"user_id": userId,
	})
}
