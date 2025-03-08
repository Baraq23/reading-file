package utils

import (
	"errors"
	"testing"
)

func TestCheckFlags(t *testing.T) {
	cases := []struct {
		name   string
		args   []string
		expect error
	}{
		{
			name:   "Valid index command",
			args:   []string{"textindex", "-c", "index", "-i", "input.txt", "-s", "1024", "-o", "index.idx"},
			expect: nil,
		},
		{
			name:   "Valid lookup command",
			args:   []string{"textindex", "-c", "lookup", "-i", "index.idx", "-q", "query_text"},
			expect: nil,
		},
		{
			name:   "Unknown flag",
			args:   []string{"textindex", "-x", "unknown"},
			expect: errors.New("flag not used"),
		},
		{
			name:   "Missing required flags",
			args:   []string{"textindex", "-c", "index"},
			expect: errors.New("usage: textindex  -c index  -i <input_file.txt>  -s <chunk-size>  -o <index_file.idx>\ntextindex  -c lookup  -i <index_file.idx> -q <query_text>"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := CheckFlags(c.args)
			if (err == nil && c.expect != nil) || (err != nil && c.expect == nil) || (err != nil && c.expect != nil && err.Error() != c.expect.Error()) {
				t.Errorf("%s: expected error %v, got %v", c.name, c.expect, err)
			}
		})
	}
}
