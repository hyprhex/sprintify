package main

import (
	"fmt"
	"os"
)

type Config struct {
	DBPort string
	DBUser string
	DBPass string
	DBAdrs string
	DBName string
	JWTSec string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		DBPort: getEnv("PORT", "5432"),
		DBUser: getEnv("DB_USER", "swap"),
		DBPass: getEnv("DB_PASS", "swap"),
		DBAdrs: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName: getEnv("DB_NAME", "sprintify"),
		JWTSec: getEnv("JWT_SECRET", "randomjwtsecret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
