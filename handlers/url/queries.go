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

func getRowWithURL(originalUrl string, s *state.State) (URLResponse, error) {
	var url URLResponse
	row := s.Database.QueryRow("SELECT id, original_url, short_token FROM urls WHERE original_url = ?", originalUrl)
	err := row.Scan(&url.ID, &url.OriginalURL, &url.NewToken)

	if err != nil {
		return url, err
	}

	return url, nil

}

func createOrUpdateRow(originalUrl string, s *state.State) (URLResponse, error) {
	var urlResponse URLResponse
	var err error

	urlResponse, err = getRowWithURL(originalUrl, s)

	if err == nil {
		return urlResponse, ErrAlreadyExistingRow
	}

	if err != sql.ErrNoRows {
		return urlResponse, err
	}

	urlResponse.OriginalURL = originalUrl
	urlResponse.NewToken, err = GenerateRandomString(6)

	if err != nil {
		return urlResponse, err
	}

	s.Database.QueryRow(
		"INSERT INTO urls (original_url, short_token) VALUES (?, ?)",
		urlResponse.OriginalURL, urlResponse.NewToken,
	)
	return urlResponse, nil
}
