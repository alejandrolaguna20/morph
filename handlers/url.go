package handlers

import (
	"net/http"

	"github.com/alejandrolaguna20/morph/handlers/url"
	"github.com/alejandrolaguna20/morph/state"
)

func SetupURLHandlers(s *state.State) {
	http.HandleFunc("/url/", func(w http.ResponseWriter, r *http.Request) {
		url.GetShortenedURLHandler(s, w, r)
	})
}
