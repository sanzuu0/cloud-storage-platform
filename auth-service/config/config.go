package config

import (
	"log"
	"os"
	"time"
)

// TODO: Реализовать загрузку конфигурации
//  - Считать .env файл
//  - Загрузить настройки базы данных, Redis, Kafka, JWT секреты и пр.
//
// Можно использовать библиотеку типа godotenv + кастомные структуры для конфига.

type Config struct {
	HTTPPort        string
	PostgresDSN     string
	RedisAddr       string
	JWTSecret       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func Load() Config {
	cfg := Config{
		HTTPPort:        getEnv("HTTP_PORT", ":8080"),
		PostgresDSN:     mustEnv("POSTGRES_DSN"),
		RedisAddr:       mustEnv("REDIS_ADDR"),
		JWTSecret:       mustEnv("JWT_SECRET"),
		AccessTokenTTL:  getEnvDuration("ACCESS_TOKEN_TTL", 15*time.Minute),
		RefreshTokenTTL: getEnvDuration("REFRESH_TOKEN_TTL", 30*24*time.Hour),
	}
	return cfg
}

func mustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("required environment variable not set: %s", key)
	}
	return val
}

func getEnv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	dur, err := time.ParseDuration(val)
	if err != nil {
		log.Printf("invalid duration for %s: %s, using default %s", key, val, def)
		return def
	}
	return dur
}
