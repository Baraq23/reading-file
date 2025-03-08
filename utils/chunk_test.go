package utils

import "testing"

func TestChunk(t *testing.T) {
	cases := []struct {
		name    string
		content []byte
		sep     int
		expect  [][]byte
	}{
		{
			name:    "Chunking 8-byte content into 4-byte chunks",
			content: []byte("abcdefgh"),
			sep:     4,
			expect:  [][]byte{[]byte("abcd"), []byte("efgh")},
		},
		{
			name:    "Chunking 10-byte content into 3-byte chunks",
			content: []byte("abcdefghij"),
			sep:     3,
			expect:  [][]byte{[]byte("abc"), []byte("def"), []byte("ghi"), []byte("j")},
		},
		{
			name:    "Chunking empty content",
			content: []byte(""),
			sep:     4,
			expect:  [][]byte{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := Chunk(c.content, c.sep)
			if len(result) != len(c.expect) {
				t.Errorf("%s: expected %d chunks, got %d", c.name, len(c.expect), len(result))
			}
			for i := range result {
				if string(result[i]) != string(c.expect[i]) {
					t.Errorf("%s: expected chunk %d to be %q, got %q", c.name, i, string(c.expect[i]), string(result[i]))
				}
			}
		})
	}
}
