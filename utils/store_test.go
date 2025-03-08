package utils

import "testing"

func TestStore(t *testing.T) {
	cases := []struct {
		name    string
		content []byte
		simHash uint64
		index   int
	}{
		{
			name:    "Store a simple entry",
			content: []byte("example content"),
			simHash: 123456789,
			index:   42,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			store := make(map[uint64]struct {
				value  []byte
				offset int
			})

			Store(store, c.content, c.simHash, c.index)

			entry, exists := store[c.simHash]
			if !exists {
				t.Errorf("%s: expected entry to be stored, but it was not", c.name)
			}
			if string(entry.value) != string(c.content) {
				t.Errorf("%s: expected stored content %q, got %q", c.name, string(c.content), string(entry.value))
			}
			if entry.offset != c.index {
				t.Errorf("%s: expected stored index %d, got %d", c.name, c.index, entry.offset)
			}
		})
	}
}
