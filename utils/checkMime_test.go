package utils

import (
	"errors"
	"mime/multipart"
	"os"
	"testing"
)

func TestCheckMime(t *testing.T) {
	cases := []struct {
		name   string
		path   string
		expect error
	}{
		{
			name:   "Valid TXT file",
			path:   "testdata/document.txt",
			expect: nil,
		},
		{
			name:   "Valid PDF file",
			path:   "testdata/report.pdf",
			expect: nil,
		},
		{
			name:   "Invalid file type (JPEG)",
			path:   "testdata/image.jpeg",
			expect: errors.New("file not allowed"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			file, err := os.Open(c.path)
			if err != nil {
				t.Fatalf("failed to open test file: %v", err)
			}
			defer file.Close()

			var multipartFile multipart.File = file
			err = CheckMime(multipartFile)
			if (err == nil && c.expect != nil) || (err != nil && c.expect == nil) || (err != nil && c.expect != nil && err.Error() != c.expect.Error()) {
				t.Errorf("%s: expected error %v, got %v", c.name, c.expect, err)
			}
		})
	}
}
