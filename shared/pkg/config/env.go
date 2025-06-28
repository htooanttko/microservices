package config

import "os"

func GetEnv(k, fb string) string {
	v := os.Getenv(k)
	if v == "" {
		return fb // fallback
	}
	return v
}
