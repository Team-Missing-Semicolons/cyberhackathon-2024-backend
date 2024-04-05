package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Server struct {
	Port               string
	CorsAllowedOrigins []string
}

type Database struct {
	DSN string
}

// Config stores the app configuration.
type Config struct {
	Server   Server
	Database Database
}

// New loads Config, using .env.template as the config source, and returns it.
func New(useDotenv bool) (*Config, error) {
	if useDotenv {
		if err := godotenv.Load(".env.template"); err != nil {
			return nil, fmt.Errorf("failed to load .env.template: %w", err)
		}
	}

	return &Config{
		Server: Server{
			Port:               os.Getenv("PORT"),
			CorsAllowedOrigins: parseOrigins(os.Getenv("CORS_ALLOWED_ORIGINS")),
		},
		Database: Database{
			DSN: os.Getenv("DSN"),
		},
	}, nil
}

func parseOrigins(origins string) []string {
	return strings.Split(origins, ",")
}
