package handlers

import "net/http"

func HandlersSetup() {
	http.HandleFunc("/hello", HelloWorldHandler)
	http.HandleFunc("/url", PostShortenUrlHandler)
}
