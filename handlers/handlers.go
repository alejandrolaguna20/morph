package handlers

import (
	"net/http"

	"github.com/alejandrolaguna20/morph/state"
)

func HandlersSetup(s *state.State) {
	http.HandleFunc("/hello", HelloWorldHandler)
	SetupURLHandlers(s)
}
