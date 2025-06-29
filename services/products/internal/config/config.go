package config

import "time"

type Config struct {
	AppEnv   string
	Server   ServerConfig
	Database DatabaseConfig
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
	SSLMode  string
}
