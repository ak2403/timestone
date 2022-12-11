package timestone

import (
	"errors"
	"strconv"
	"time"
)

func extractYear(now time.Time, format string) (string, error) {
	year := strconv.Itoa(now.Year())

	switch format {
	case "yyyy":
		return year, nil
	case "yyy":
		return year[1:], nil
	case "yy":
		return year[2:], nil
	default:
		return "", errors.New("Unsupported year format")
	}
}

func Year(format string) (string, error) {
	now := time.Now()

	return extractYear(now, format)
}
