package config

type JwtAuthConfig struct {
	AccessSecretKey  string `env:"JWT_ACCESS_SECRET_KEY"`
	RefreshSecretKey string `env:"JWT_REFRESH_SECRET_KEY"`
}
