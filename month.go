package timestone

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func extractMonth(now time.Time, format string) (string, error) {
	month := fmt.Sprintf("%s", now.Month())

	switch format {
	case "MM":
		return month, nil
	case "mm":
		return strings.ToLower(month), nil
	case "M":
		return month[:3], nil
	case "m":
		return strings.ToLower(month[:3]), nil
	default:
		return "", errors.New("Unsupported month format")
	}
}

func Month(format string) (string, error) {
	now := time.Now()

	return extractMonth(now, format)
}
