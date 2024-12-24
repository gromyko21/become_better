package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	database "become_better/src/db"
)

type CommonConfig struct {
	Host     string
	GRPcPort string
	HTTPport string
}

type App struct {
	Postgres *database.Postgres
}

type PostgresConfig struct {
	ConnString string
}

type Config struct {
	CommonConfig
	PostgresConfig
}

// New returns a new Config struct
func New() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		logrus.Fatalf("Error loading .env file %v", err)
	}
	return &Config{
		CommonConfig: CommonConfig{
			Host:     getEnv("HOST", ""),
			GRPcPort: getEnv("GRPC_PORT", ""),
			HTTPport: getEnv("HTTP_PORT", ""),
		},
		PostgresConfig: PostgresConfig{
			ConnString: getEnv("PGConnString", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
