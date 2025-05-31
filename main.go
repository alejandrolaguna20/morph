package main

import (
	"net/http"

	"github.com/alejandrolaguna20/morph/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
