package service

import "twitter-go/internal/service/interfaces"

type Service struct {
	AuthService interfaces.AuthService
	PostService interfaces.PostService
}
