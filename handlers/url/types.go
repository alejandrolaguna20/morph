package url

import "errors"

type URLResponse struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"url"`
	NewToken    string `json:"short_token"`
}

type PostURLRequest struct {
	URL string `json:"url"`
}

var (
	ErrInvalidURL         = errors.New("invalid URL format: not enough path segments")
	ErrNotANumericID      = errors.New("not a numeric ID")
	ErrNotFound           = errors.New("URL not found")
	ErrAlreadyExistingRow = errors.New("Already existing row")
)
