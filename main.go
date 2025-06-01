package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/alejandrolaguna20/morph/handlers"
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

func handlersSetup() {
	http.HandleFunc("/hello", handlers.HelloWorldHandler)
	http.HandleFunc("/url", handlers.PostShortenUrlHandler)
}

func getEnvOrFatal(s string) string {
	value := os.Getenv(s)
	if value == "" {
		log.Fatalf("Environment variable [%s] not set", s)
	}
	return value
}

func loadEnv() Env {
	env := Env{}
	err := dotenv.Load()

	if err != nil {
		log.Fatal(".env file not loaded, aborting")
	}

	env.DatabaseHost = getEnvOrFatal("DB_HOST")
	env.DatabasePassword = getEnvOrFatal("DB_PASSWORD")
	env.DatabaseName = getEnvOrFatal("DB_NAME")
	env.DatabaseUser = getEnvOrFatal("DB_USER")

	port, err := strconv.Atoi(getEnvOrFatal("DB_PORT"))
	if err != nil {
		log.Fatal("Database port missing")
	} else {
		env.DatabasePort = port
	}

	port, err = strconv.Atoi(getEnvOrFatal("PORT"))
	if err != nil {
		log.Fatal("Server port missing")
	} else {
		env.ServerPort = port
	}

	return env
}

func main() {
	env := loadEnv()
	portString := ":" + strconv.Itoa(env.ServerPort)
	handlersSetup()
	log.Fatal(http.ListenAndServe(portString, nil))
}
