package utils

import (
	"errors"
	"testing"
)

func TestDecodeBinary(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		expect    map[uint64]Entry
		expectErr error
	}{
		{
			name: "Valid idx file",
			path: "testdata/index.idx",
			expect: map[uint64]Entry{
				123456789: {[]byte("example content"), 42},
			},
			expectErr: nil,
		},
		{
			name:      "Invalid file type",
			path:      "testdata/document.txt",
			expect:    nil,
			expectErr: errors.New("wrong type of file"),
		},
		{
			name:      "Corrupted idx file",
			path:      "testdata/corrupt.idx",
			expect:    nil,
			expectErr: errors.New("error while decoding file"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := DecodeBinary(c.path)

			if (err == nil && c.expectErr != nil) || (err != nil && c.expectErr == nil) || (err != nil && c.expectErr != nil && err.Error() != c.expectErr.Error()) {
				t.Errorf("%s: expected error %v, got %v", c.name, c.expectErr, err)
			}

			if c.expectErr == nil {
				if len(result) != len(c.expect) {
					t.Errorf("%s: expected map size %d, got %d", c.name, len(c.expect), len(result))
				}
				for key, val := range c.expect {
					resVal, exists := result[key]
					if !exists {
						t.Errorf("%s: expected key %d to exist, but it does not", c.name, key)
					}
					if string(resVal.Value) != string(val.Value) {
						t.Errorf("%s: expected value %q, got %q", c.name, string(val.Value), string(resVal.Value))
					}
					if resVal.Offset != val.Offset {
						t.Errorf("%s: expected offset %d, got %d", c.name, val.Offset, resVal.Offset)
					}
				}
			}
		})
	}
}
