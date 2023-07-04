package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserPayload struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type JwtResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func GenerateJwt(payload UserPayload) (*JwtResponse, error) {
	accessTokenKey := []byte("secret")
	refreshTokenKey := []byte("super-secret")

	accessTokenClaims := jwt.MapClaims{
		"payload": payload,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	refreshTokenClaims := jwt.MapClaims{
		"userId": payload.Id,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	access, err := accessToken.SignedString(accessTokenKey)
	if err != nil {
		return nil, err
	}

	refresh, err := refreshToken.SignedString(refreshTokenKey)
	if err != nil {
		return nil, err
	}

	return &JwtResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}
