package utils

import "testing"

func TestSimHash(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect uint64
	}{
		{
			name:   "SimHash of ['hello', 'world']",
			input:  []string{"hello", "world"},
			expect: SimHash([]string{"hello", "world"}),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := SimHash(c.input)
			if result != c.expect {
				t.Errorf("%s: expected %d, got %d", c.name, c.expect, result)
			}
		})
	}
}
