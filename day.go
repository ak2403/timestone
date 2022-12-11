package timestone

import (
	"strconv"
	"time"
)

func extractDay(now time.Time, format string) string {
	day := strconv.Itoa(now.Day())

	return day
}

func Day(format string) string {
	now := time.Now()

	return extractDay(now, format)
}
