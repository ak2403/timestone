package timestone

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type MonthFormat int

const (
	MM MonthFormat = iota
	M
	mm
	m
)

func (m MonthFormat) toString() string {
	return []string{"MM", "M", "mm", "m"}[m]
}

func extractMonth(now time.Time, format string) (string, error) {
	month := strconv.Itoa(int(now.Month()))

	switch format {
	case MM.toString():
		return month, nil
	case mm.toString():
		return strings.ToLower(month), nil
	case M.toString():
		return month[:3], nil
	case m.toString():
		return strings.ToLower(month[:3]), nil
	default:
		return "", errors.New("Unsupported month format")
	}
}

func Month(format string) (string, error) {
	now := time.Now()

	return extractMonth(now, format)
}
