package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel        string
	PBAdminEmail    string
	PBAdminPassword string
	PBDebug         bool
	ClientID        string
	ClientSecret    string
	RedirectURI     string
	AuthURI         string
	TokenURI        string
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
		AuthURI:         os.Getenv("AUTH_URI"),
		TokenURI:        os.Getenv("TOKEN_URI"),
	}, nil
}
