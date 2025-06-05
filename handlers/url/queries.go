package url

import (
	"database/sql"

	"github.com/alejandrolaguna20/morph/state"
)

func findURLWithToken(token string, s *state.State) (URLResponse, error) {
	var url URLResponse
	row := s.Database.QueryRow("SELECT id, original_url, short_token FROM urls WHERE short_token = ?", token)
	err := row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return url, ErrNotFound
		}
		return url, err
	}
	return url, nil
}

func findURLWithID(id int, s *state.State) (URLResponse, error) {
	var url URLResponse
	row := s.Database.QueryRow("SELECT id, original_url, short_token FROM urls WHERE id = ?", id)
	err := row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return url, ErrNotFound
		}
		return url, err
	}
	return url, nil
}
