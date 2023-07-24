package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel string
	AppMode  string

	PBAdminEmail    string
	PBAdminPassword string
	PBDebug         bool

	ClientID     string
	ClientSecret string
	RedirectURI  string
	BaseURL      string
	AuthURI      string
	TokenURI     string

	SecretKey string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		PBDebug:         os.Getenv("PB_DEBUG") == "true",
		LogLevel:        os.Getenv("LOG_LEVEL"),
		PBAdminEmail:    os.Getenv("PB_ADMIN_EMAIL"),
		PBAdminPassword: os.Getenv("PB_ADMIN_PASSWORD"),
		ClientID:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		RedirectURI:     os.Getenv("REDIRECT_URI"),
		BaseURL:         os.Getenv("BASE_URL"),
		AuthURI:         os.Getenv("AUTH_URI"),
		TokenURI:        os.Getenv("TOKEN_URI"),
		AppMode:         os.Getenv("APP_MODE"),
		SecretKey:       os.Getenv("SECRET_KEY"),
	}, nil
}
