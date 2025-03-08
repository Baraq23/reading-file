package utils

import (
	"encoding/gob"
	"os"
	"testing"
)

func TestStoreBinary(t *testing.T) {
	cases := []struct {
		name  string
		store map[uint64]struct {
			value  []byte
			offset int
		}
		filePath string
	}{
		{
			name: "Store and serialize map",
			store: map[uint64]struct {
				value  []byte
				offset int
			}{
				123456789: {value: []byte("example content"), offset: 42},
			},
			filePath: "testdata/index.idx",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			file, err := os.Create(c.filePath)
			if err != nil {
				t.Fatalf("failed to create file: %v", err)
			}
			defer os.Remove(c.filePath)
			defer file.Close()

			err = StoreBinary(c.store, file)
			if err != nil {
				t.Errorf("%s: unexpected error: %v", c.name, err)
			}

			// Verify file is created
			_, err = os.Stat(c.filePath)
			if os.IsNotExist(err) {
				t.Errorf("%s: expected file to be created, but it was not", c.name)
			}

			// Verify stored data
			file, err = os.Open(c.filePath)
			if err != nil {
				t.Fatalf("failed to open stored file: %v", err)
			}
			defer file.Close()

			decoder := gob.NewDecoder(file)
			var loadedStore map[uint64]struct {
				value  []byte
				offset int
			}
			if err := decoder.Decode(&loadedStore); err != nil {
				t.Fatalf("failed to decode stored data: %v", err)
			}

			if len(loadedStore) != len(c.store) {
				t.Errorf("%s: expected map size %d, got %d", c.name, len(c.store), len(loadedStore))
			}
		})
	}
}
