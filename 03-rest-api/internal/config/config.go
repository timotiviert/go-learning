package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	LogLevelInfo  = "INFO"
	LogLevelDebug = "DEBUG"
	LogLevelError = "ERROR"
	LogLevelFatal = "FATAL"
	LogLevelPanic = "PANIC"
)

type Config struct {
	SeverPort   int
	DatabaseDSN string
	JWTSecret   string
	LogLevel    string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Maybe Get per property makes sense, as almost all need some parsing (i.e. port to int etc.)
	pStr := getEnv("SERVER_PORT", "8080")
	sp, err := strconv.Atoi(pStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT %s: %w", sp, err)
	}

	dsn, err := getDatabaseDSN()
	if err != nil {
		return nil, fmt.Errorf("failed to get database DSN: %w", err)
	}

	jwt := os.Getenv("JWT_SECRET")
	if jwt == "" {
		return nil, fmt.Errorf("required environment variable JWT_SECRET is missing")
	}

	lvl := getEnv("LOG_LEVEL", LogLevelInfo)

	return &Config{
		SeverPort:   sp,
		DatabaseDSN: dsn,
		JWTSecret:   jwt,
		LogLevel:    lvl,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getDatabaseDSN() (string, error) {
	return "", nil
}
