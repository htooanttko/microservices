package config

import (
	"time"

	"github.com/htooanttko/microservices/shared/pkg/config"
	"github.com/htooanttko/microservices/shared/pkg/logger"
	"github.com/joho/godotenv"
)

func LoadConfig(p string) (*Config, error) {
	if err := godotenv.Load(p); err != nil {
		logger.Error.Printf("No .env file found at %s.", p)
	}

	// Parse JWT expiration duration
	jwtExpiration, err := time.ParseDuration(config.GetEnv("JWT_EXPIRATION", "24h"))
	if err != nil {
		return nil, err
	}

	// Parse server timeouts
	readTimeout, err := time.ParseDuration(config.GetEnv("SERVER_TIMEOUT_READ", "5s"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := time.ParseDuration(config.GetEnv("SERVER_TIMEOUT_WRITE", "10s"))
	if err != nil {
		return nil, err
	}

	idleTimeout, err := time.ParseDuration(config.GetEnv("SERVER_TIMEOUT_IDLE", "15s"))
	if err != nil {
		return nil, err
	}

	return &Config{
		AppEnv: config.GetEnv("APP_ENV", "development"),
		Server: ServerConfig{
			Port:         config.GetEnv("SERVER_PORT", "9090"),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
		Database: DatabaseConfig{
			Host:     config.GetEnv("DB_HOST", "localhost"),
			Port:     config.GetEnv("DB_PORT", "5432"),
			User:     config.GetEnv("DB_USER", ""),
			Password: config.GetEnv("DB_PASSWORD", ""),
			Name:     config.GetEnv("DB_NAME", ""),
			SSLMode:  config.GetEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:     config.GetEnv("JWT_SECRET", ""),
			Expiration: jwtExpiration,
		},
		LogLevel: config.GetEnv("LOG_LEVEL", "debug"),
	}, nil
}
