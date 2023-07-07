package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"twitter-go/entity"
	"twitter-go/internal/service/interfaces"
	"twitter-go/utils"
)

type (
	PostsHandler interface {
		Create(c echo.Context) error
	}

	postsHandler struct {
		PostsService interfaces.PostService
	}
)

func (h *postsHandler) Create(c echo.Context) error {
	var req entity.CreatePostRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user := c.Get("user").(utils.UserPayload)

	postId, err := h.PostsService.CreateNewPost(req, user.Id)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"postId": postId,
	})
}
