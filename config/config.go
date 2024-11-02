package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Port string `envconfig:"APP_PORT" required:"true"`
		Host string `envconfig:"APP_HOST" required:"true"`
	}
	Database struct {
		StringConn string `envconfig:"DB_STRING_CONN" required:"true"`
		MinConn    string `envconfig:"DB_MIN_CONN" required:"true"`
		MaxConn    string `envconfig:"DB_MAX_CONN" required:"true"`
	}
	Jwt struct {
		AccessTokenSecret           string `envconfig:"JWT_ACCESS_TOKEN_SECRET" required:"true"`
		AccessTokenExpirationInHour int    `envconfig:"JWT_ACCESS_TOKEN_EXPIRATION_IN_HOUR" required:"true"`
	}
}

func LoadConfig(fileName string) *Config {
	var config Config
	err := godotenv.Overload(fileName)
	if err != nil {
		log.Fatalf("Failed to overload env vars: %v", err)
	}
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Failed to populates env vars: %v", err)
	}
	return &config
}
