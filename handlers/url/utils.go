package url

import "strconv"

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
