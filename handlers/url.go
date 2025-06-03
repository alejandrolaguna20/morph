package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alejandrolaguna20/morph/state"
)

type URLResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"url"`
	NewToken    string `json:"new_token"`
}

func SetupURLHandlers(s *state.State) {
	http.HandleFunc("/url/", func(w http.ResponseWriter, r *http.Request) {
		postShortenUrlHandler(s, w, r)
	})
}

func getIDFromURL(path string) (int, error) {
	path = strings.Trim(path, "/")
	pathSections := strings.Split(path, "/")

	if len(pathSections) < 2 {
		return 0, errors.New("invalid URL format: not enough path segments")
	}

	idStr := pathSections[len(pathSections)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %w", err)
	}
	return id, nil
}

func postShortenUrlHandler(s *state.State, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var url URLResponse

	id, err := getIDFromURL(r.URL.Path)

	if err != nil {
		http.Error(w, "Wrong request path format", http.StatusBadRequest)
	}

	row := s.Database.QueryRow("SELECT id, original_url, new_token FROM urls WHERE id = ?", id)
	err = row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}
