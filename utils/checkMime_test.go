package utils

import (
	"os"
	"testing"
)

func TestCheckMime(t *testing.T) {
	cases := []struct {
		name   string
		path   string
		expect bool
	}{
		{
			name:   "Valid TXT file",
			path:   "testdata/document.txt",
			expect: true,
		},
		{
			name:   "Valid PDF file",
			path:   "testdata/report.pdf",
			expect: true,
		},
		{
			name:   "Invalid file type (JPEG)",
			path:   "testdata/image.jpeg",
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			file, err := os.ReadFile(c.path)
			if err != nil {
				t.Fatalf("failed to open test file: %v", err)
			}

			check := CheckMime(file)
			if (!check && c.expect) || (check && !c.expect) {
				t.Errorf("%s: expected %v, got %v", c.name, c.expect, err)
			}
		})
	}
}
