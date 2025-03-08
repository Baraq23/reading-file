package utils

import (
	"hash/fnv"
	"testing"
)

func TestHashToken(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		expect uint64
	}{
		{
			name:   "Hash of 'hello'",
			input:  "hello",
			expect: fnv.New64a().Sum64(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := HashToken(c.input)
			if result != c.expect {
				t.Errorf("%s: expected %d, got %d", c.name, c.expect, result)
			}
		})
	}
}
