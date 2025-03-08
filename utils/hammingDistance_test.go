package utils

import "testing"

func TestHammingDistance(t *testing.T) {
	cases := []struct {
		name   string
		a, b   uint64
		expect int
	}{
		{
			name:   "Identical numbers",
			a:      0b101010,
			b:      0b101010,
			expect: 0,
		},
		{
			name:   "One bit difference",
			a:      0b101010,
			b:      0b101011,
			expect: 1,
		},
		{
			name:   "Multiple bit difference",
			a:      0b101010,
			b:      0b110011,
			expect: 3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := HammingDistance(c.a, c.b)
			if result != c.expect {
				t.Errorf("%s: expected %d, got %d", c.name, c.expect, result)
			}
		})
	}
}
