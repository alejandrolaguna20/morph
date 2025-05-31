package main

import (
	"log"
	"net/http"

	"github.com/alejandrolaguna20/morph/handlers"
)

func handlersSetup() {
	http.HandleFunc("/hello", handlers.HelloWorldHandler)
	http.HandleFunc("/url", handlers.PostShortenUrlHandler)
}

func main() {
	handlersSetup()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
