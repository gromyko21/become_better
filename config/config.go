package config

import (
    "os"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type CommonConfig struct {
    Host string
    GRPcPort   string
	HTTPport string
}

type Config struct {
    CommonConfig CommonConfig
}

// New returns a new Config struct
func New() *Config {
	err := godotenv.Load("config/.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
		fmt.Println(err)
	}
    return &Config{
        CommonConfig: CommonConfig{
			Host: getEnv("HOST", ""),
			GRPcPort: getEnv("GRPC_PORT", ""),
			HTTPport: getEnv("HTTP_PORT", ""),
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
