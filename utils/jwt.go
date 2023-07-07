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

type JwtAccessClaims struct {
	jwt.RegisteredClaims
}
type JwtRefreshClaims struct {
	access string
	jwt.RegisteredClaims
}

type JwtResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func GetAccessToken(payload UserPayload, secretKey []byte) (string, error) {
	accessTokenClaims := jwt.MapClaims{
		"id":         payload.Id,
		"name":       payload.Name,
		"email":      payload.Email,
		"created_at": payload.CreatedAt,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	access, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return access, nil
}

func GetRefreshToken(access string, secretKey []byte) (string, error) {
	refreshTokenClaims := jwt.MapClaims{
		"access": access,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refresh, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refresh, nil
}

func GenerateJwt(payload UserPayload) (*JwtResponse, error) {

	jwtConfig := &config.JwtAuthConfig{}
	if err := env.Parse(jwtConfig); err != nil {
		return nil, err
	}

	accessTokenKey := []byte(jwtConfig.AccessSecretKey)
	refreshTokenKey := []byte(jwtConfig.RefreshSecretKey)

	access, err := GetAccessToken(payload, accessTokenKey)
	if err != nil {
		return nil, err
	}
	refresh, err := GetRefreshToken(access, refreshTokenKey)
	if err != nil {
		return nil, err
	}

	return &JwtResponse{
		Access:  access,
		Refresh: refresh,
	}, nil
}
