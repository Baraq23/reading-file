package utils

import (
	"errors"
	"testing"
)

func TestReadFile(t *testing.T) {
	cases := []struct {
		name          string
		path          string
		expectErr     error
		expectContent string
	}{
		{
			name:          "Valid file",
			path:          "testdata/document.txt",
			expectErr:     nil,
			expectContent: "Report for file indexing",
		},
		{
			name:          "Non-existent file",
			path:          "testdata/missing.txt",
			expectErr:     errors.New("error reading content of the file"),
			expectContent: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, err := ReadFile(c.path)
			if (err == nil && c.expectErr != nil) || (err != nil && c.expectErr == nil) || (err != nil && c.expectErr != nil && err.Error() != c.expectErr.Error()) {
				t.Errorf("%s: expected error %v, got %v", c.name, c.expectErr, err)
			}
			if err == nil && string(data) != c.expectContent {
				t.Errorf("%s: expected content %q, got %q", c.name, c.expectContent, string(data))
			}
		})
	}
}
