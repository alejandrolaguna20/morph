package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
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

var (
	ErrInvalidURL    = errors.New("invalid URL format: not enough path segments")
	ErrNotANumericID = errors.New("Not a numeric ID")
	ErrNotFound      = errors.New("URL not found")
)

func isInteger(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func SetupURLHandlers(s *state.State) {
	http.HandleFunc("/url/", func(w http.ResponseWriter, r *http.Request) {
		getShortenedURLHandler(s, w, r)
	})
}

func getIDFromURL(idStr string) (int, error) {

	if !isInteger(idStr) {
		return 0, ErrNotANumericID
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, ErrNotANumericID
	}
	return id, nil
}

func findURLWithToken(token string, s *state.State) (URLResponse, error) {

	var url URLResponse
	row := s.Database.QueryRow("SELECT id, original_url, new_token FROM urls WHERE new_token = ?", token)
	err := row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)

	if err != nil {
		err = ErrNotFound
	}

	return url, err
}

func findURLWithID(id int, s *state.State) (URLResponse, error) {
	var url URLResponse
	row := s.Database.QueryRow("SELECT id, original_url, new_token FROM urls WHERE id = ?", id)
	err := row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)

	if err != nil {
		err = ErrNotFound
	}

	return url, err
}

func getShortenedURLHandler(s *state.State, w http.ResponseWriter, r *http.Request) {

	var searchByToken bool = false
	var url URLResponse
	var err error

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	path := strings.Trim(r.URL.Path, "/")
	pathSections := strings.Split(path, "/")

	if len(pathSections) < 2 {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	idStr := pathSections[len(pathSections)-1]

	id, err := getIDFromURL(idStr)

	if err != nil {
		searchByToken = true
	}

	if searchByToken {
		url, err = findURLWithToken(idStr, s)
	} else {
		url, err = findURLWithID(id, s)
	}

	if err != nil {
		if err == sql.ErrNoRows || err == ErrNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}
