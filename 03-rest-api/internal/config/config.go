package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type LogLevel string

const (
	LogLevelTrace = "TRACE"
	LogLevelDebug = "DEBUG"
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
	LogLevelFatal = "FATAL"
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelTrace, LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError, LogLevelFatal:
		return true
	}
	return false
}

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

	lvl := strings.ToUpper(getEnv("LOG_LEVEL", LogLevelInfo))
	if !LogLevel(lvl).IsValid() {
		return nil, fmt.Errorf("invalid LOG_LEVEL %s", lvl)
	}

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
