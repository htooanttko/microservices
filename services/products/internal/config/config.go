package config

import "time"

type Config struct {
	AppEnv   string
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	LogLevel string
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}
