package state

import (
	"log"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

type Env struct {
	DatabaseName     string
	DatabasePassword string
	DatabaseUser     string
	DatabasePort     int
	DatabaseHost     string
	ServerPort       int
}

func getEnvOrFatal(s string) string {
	value := os.Getenv(s)
	if value == "" {
		log.Fatalf("[ERROR] Environment variable [%s] not set", s)
	}
	return value
}

func getEnvIntOrFatal(s string) int {
	value, err := strconv.Atoi(getEnvOrFatal(s))
	if err != nil {
		log.Fatalf("[ERROR] Environment variable [%s] cannot be converted to int", s)
	}
	return value
}

func loadEnv() Env {
	var env Env
	err := dotenv.Load()

	if err != nil {
		log.Fatal("[ERROR] .env file not loaded, aborting")
	}

	env.DatabaseHost = getEnvOrFatal("DB_HOST")
	env.DatabasePassword = getEnvOrFatal("DB_PASSWORD")
	env.DatabaseName = getEnvOrFatal("DB_NAME")
	env.DatabaseUser = getEnvOrFatal("DB_USER")
	env.DatabasePort = getEnvIntOrFatal("DB_PORT")
	env.ServerPort = getEnvIntOrFatal("PORT")

	return env
}
