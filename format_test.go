package timestone

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {

	type formatTest struct {
		date   time.Time
		format string
		output string
	}

	tests := []formatTest{
		{date: time.Now(), format: "dd-M-yyyy", output: "11-Dec-2022"},
		{date: time.Now(), format: "dd-MM-yyyy", output: "11-December-2022"},
		{date: time.Now(), format: "dd-m-yyy", output: "11-dec-022"},
		{date: time.Now(), format: "dd-mm-yy", output: "11-december-22"},
		{date: time.Now(), format: "M-yy", output: "Dec-22"},
		{date: time.Now(), format: "mm-yyy", output: "december-022"},
	}

	for _, v := range tests {
		format := Format(v.date, v.format)

		if format != v.output {
			t.Error("Expected ", v.output, "but got ", format)
		}
	}
}
