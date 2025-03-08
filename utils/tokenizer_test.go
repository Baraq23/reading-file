package utils

import "testing"

func TestTokenizer(t *testing.T) {
	cases := []struct {
		name    string
		content []byte
		expect  []string
	}{
		{
			name:    "Simple sentence",
			content: []byte("Hello there you"),
			expect:  []string{"Hello", "there", "you"},
		},
		{
			name:    "Multiple spaces",
			content: []byte("  This  is   a test  "),
			expect:  []string{"This", "is", "a", "test"},
		},
		{
			name:    "Empty string",
			content: []byte(""),
			expect:  []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Tokenizer(c.content)
			if len(result) != len(c.expect) {
				t.Errorf("%s: expected %v, got %v", c.name, c.expect, result)
			}
			for i := range result {
				if result[i] != c.expect[i] {
					t.Errorf("%s: expected word %d to be %q, got %q", c.name, i, c.expect[i], result[i])
				}
			}
		})
	}
}
