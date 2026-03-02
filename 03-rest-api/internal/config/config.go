package config

import (
	"log"
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
		log.Fatalf("failed to load .env file: %v", err)
	}

	// Maybe Get per property makes sense, as almost all need some parsing (i.e. port to int etc.)

	sp, err := strconv.Atoi(getEnv("SEVER_PORT", "8080"))

	if err != nil {
		log.Fatalf("failed to parse sever port: %v", err)
	}

	dsn, err := getDatabaseDSN()
	if err != nil {
		log.Fatalf("failed to get database DSN: %v", err)
	}

	return &Config{
		SeverPort:   sp,
		DatabaseDSN: dsn,
		JWTSecret:   getEnvRequired("JWT_SECRET"),
		LogLevel:    getEnv("LOG_LEVEL", LogLevelInfo),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	log.Fatalf("required environment variable %s not set", key)
}

func getDatabaseDSN() (string, error) {
	return "", nil
}
