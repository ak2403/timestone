package timestone

import (
	"strings"
	"time"
)

const delimiter string = "-"

func uniqueString(s string) string {
	unique := map[string]bool{}

	for _, v := range strings.Split(strings.ToLower(s), "") {
		if unique[v] {
			continue
		}
		unique[v] = true
	}

	keys := make([]string, 0, len(unique))
	for k := range unique {
		keys = append(keys, k)
	}

	return strings.Join(keys, "")
}

func extractDateAttr(format string) string {
	switch uniqueString(format) {
	case "m":
		month, _ := Month(format)

		return month
	case "y":
		year, _ := Year(format)

		return year
	case "d":
		day := Day(format)

		return day
	default:
		return ""
	}
}

func Format(date time.Time, format string) string {
	formatArr := strings.Split(format, delimiter)

	formattedDateArr := []string{}

	for _, v := range formatArr {

		formattedDateArr = append(formattedDateArr, extractDateAttr(v))
	}

	return strings.Join(formattedDateArr, delimiter)
}
