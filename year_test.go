package timestone

import (
	"testing"
)

func TestYear(t *testing.T) {

	type yearTest struct {
		format string
		output string
	}

	tests := []yearTest{
		{format: "yyyy", output: "2022"},
		{format: "yyy", output: "022"},
		{format: "yy", output: "22"},
	}

	for _, v := range tests {
		year, _ := Year(v.format)

		if year != v.output {
			t.Error("Expected : ", v.output, "but got ", year)
		}
	}
}
