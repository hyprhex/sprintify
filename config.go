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
		DBPort: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", "root"),
		DBAdrs: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "172.17.0.2"), getEnv("DB_PORT", "3306")),
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
