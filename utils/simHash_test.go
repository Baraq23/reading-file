package utils

import "testing"

func TestSimHash(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		expect uint64
	}{
		{
			name:   "SimHash of ['hello', 'world' ']",
			input:  []string{"hello", "world"},
			expect: 292971770938951683,
		},
		{
			name:   "SimHash of an empty list",
			input:  []string{},
			expect: 0,
		},
		{
			name:   "SimHash of a single word",
			input:  []string{"golang"},
			expect: 4929537852214148147,
		},
		{
			name:   "SimHash of ['foo', 'bar', 'baz']",
			input:  []string{"foo", "bar", "baz"},
			expect: 16101355973858322,
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
