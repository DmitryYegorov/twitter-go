package utils

import (
	"github.com/caarlos0/env/v9"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"twitter-go/config"
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

	jwtConfig := &config.JwtAuthConfig{}
	if err := env.Parse(jwtConfig); err != nil {
		return nil, err
	}

	accessTokenKey := []byte(jwtConfig.AccessSecretKey)
	refreshTokenKey := []byte(jwtConfig.RefreshSecretKey)

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
