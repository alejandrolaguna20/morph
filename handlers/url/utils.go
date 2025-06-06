package url

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
)

func isInteger(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
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

// - base64 encodes 6 bits per character
// - 6 bytes (6 * 8 = 48 bits) -> 8 base64 characters (48 / 6 = 8)
func GenerateRandomString(nBytes int) (string, error) {
	bytes := make([]byte, nBytes)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// wont have = + or -
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
