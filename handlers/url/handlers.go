package url

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/alejandrolaguna20/morph/state"
)

func GetShortenedURLHandler(s *state.State, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.Trim(r.URL.Path, "/")
	pathSections := strings.Split(path, "/")

	if len(pathSections) < 2 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	idStr := pathSections[len(pathSections)-1]
	id, err := getIDFromURL(idStr)

	var url URLResponse
	if err != nil {
		url, err = findURLWithToken(idStr, s)
	} else {
		url, err = findURLWithID(id, s)
	}

	if err != nil {
		if err == ErrNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func PostShortenURL(s *state.State, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody PostURLRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "URL is missing", http.StatusBadRequest)
		return
	}

	urlResponse, err := createOrUpdateRow(requestBody.URL, s)

	if err != nil && !errors.Is(err, ErrAlreadyExistingRow) {
		http.Error(w, "Could not store the URL to the database", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")

	if errors.Is(err, ErrAlreadyExistingRow) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(urlResponse)
}
