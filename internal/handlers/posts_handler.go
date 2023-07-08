package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"twitter-go/entity"
	"twitter-go/internal/service/interfaces"
	"twitter-go/utils"
)

type (
	PostsHandler interface {
		Create(c echo.Context) error
		GetAllByUser(c echo.Context) error
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"postId": postId,
	})
}

func (h *postsHandler) GetAllByUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must by a number")
	}

	logrus.Printf("Invoked GetAllByUser (%d)", userId)

	list, err := h.PostsService.GetUserPosts(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, list)
}
