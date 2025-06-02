package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alejandrolaguna20/morph/handlers"
	_ "github.com/go-sql-driver/mysql"
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

func connectToDatabase(env Env) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
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
