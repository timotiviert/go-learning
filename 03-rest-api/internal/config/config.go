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
	SeverPort    int
	DBConnString string
	JWTSecret    string
	LogLevel     string
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

	connString, err := getPGConnString()
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
		SeverPort:    sp,
		DBConnString: connString,
		JWTSecret:    jwt,
		LogLevel:     lvl,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getPGConnString() (string, error) {
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")
	user := os.Getenv("PG_USER")
	if user == "" {
		return "", fmt.Errorf("required environment variable PG_USER is missing")
	}
	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		return "", fmt.Errorf("required environment variable PG_PASSWORD is missing")
	}
	database := os.Getenv("PG_DATABASE")
	if database == "" {
		return "", fmt.Errorf("required environment variable PG_DATABASE is missing")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database), nil
}
