package handlers

import (
	"encoding/json"
	"net/http"
)

type URLResponse struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func PostShortenUrlHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(URLResponse{ID: 2, URL: "/test"})
}
