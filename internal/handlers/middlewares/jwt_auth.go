package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
	"twitter-go/utils"
)

func JwtAuth(accessSecretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get("Authorization")

			if authorization == "" {
				logrus.Printf(`User are not authorized`)
				return echo.NewHTTPError(http.StatusUnauthorized, "you are not authorized")
			}

			tokenString := strings.Split(authorization, " ")[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusInternalServerError, "unexpected signing method")
				}

				return []byte(accessSecretKey), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			claims := token.Claims.(jwt.MapClaims)
			user := utils.UserPayload{
				Id:        int(claims["id"].(float64)),
				Name:      claims["name"].(string),
				Email:     claims["email"].(string),
				CreatedAt: time.Unix(0, 0),
			}

			createdInt, err := strconv.ParseInt(claims["created_at"].(string), 10, 64)

			if err != nil {
				user.CreatedAt = time.Unix(createdInt, 0)
			}

			c.Set("user", user)

			return next(c)
		}
	}
}
