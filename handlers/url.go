package handlers

import (
	"net/http"

	"github.com/alejandrolaguna20/morph/handlers/url"
	"github.com/alejandrolaguna20/morph/state"
)

func SetupURLHandlers(s *state.State) {

	// GET /url/{id} or GET /url/{token}
	http.HandleFunc("/url/", func(w http.ResponseWriter, r *http.Request) {
		url.GetShortenedURLHandler(s, w, r)
	})

	// POST /url
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		url.PostShortenURL(s, w, r)
	})
}
