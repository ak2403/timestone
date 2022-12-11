package timestone

import (
	"testing"
	"time"
)

func TestMonth(t *testing.T) {

	type monthTest struct {
		now    time.Time
		format string
		output string
	}

	tests := []monthTest{
		{time.Now(), "M", "Dec"},
		{time.Now(), "m", "dec"},
		{time.Now(), "MM", "December"},
		{time.Now(), "mm", "december"},
	}

	for _, v := range tests {
		month, _ := extractMonth(v.now, v.format)

		if month != v.output {
			t.Error("Expected : ", v.output, "but got ", month)
		}
	}

}
