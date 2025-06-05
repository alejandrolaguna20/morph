package url

import (
	"encoding/json"
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
