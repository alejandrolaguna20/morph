package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, morph!\n")
}

func main() {
	fmt.Println("hello, morph")
	http.HandleFunc("/test", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
