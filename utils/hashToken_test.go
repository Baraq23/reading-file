package utils

import (
	"hash/fnv"
	"testing"
)

func TestHashToken(t *testing.T) {
	cases := []struct {
		name  string
		input string
	}{
		{
			name:  "Hash of 'hello'",
			input: "hello",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			val := fnv.New64a()
			val.Write([]byte(c.input))
			expect := val.Sum64()
			result := HashToken(c.input)
			if result != expect {
				t.Errorf("%s: expected %d, got %d", c.name, expect, result)
			}
		})
	}
}
