package config

type DataBaseConfig struct {
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Port     uint16 `env:"DB_PORT"`
}
