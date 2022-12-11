package timestone

import (
	"errors"
	"strconv"
	"time"
)

type YearFormat int

const (
	yyyy YearFormat = iota
	yyy
	yy
)

func (y YearFormat) toString() string {
	return []string{"yyyy", "yyy", "yy"}[y]
}

func extractYear(now time.Time, format string) (string, error) {
	year := strconv.Itoa(now.Year())

	switch format {
	case yyyy.toString():
		return year, nil
	case yyy.toString():
		return year[1:], nil
	case yy.toString():
		return year[2:], nil
	default:
		return "", errors.New("Unsupported year format")
	}
}

func Year(format string) (string, error) {
	now := time.Now()

	return extractYear(now, format)
}
