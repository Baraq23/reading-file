package utils

import (
	"encoding/gob"
	"os"
	"testing"
)
type Data struct {
	Name  string
	Store map[uint64]Entry
	FilePath string
}

var Cases = []Data {
	{
		Name: "Store and serialize map",
		Store: map[uint64]Entry{
			123456789: {Value: []byte("example content"), Offset: 42},
		},
		FilePath: "testdata/index.idx",
	},

}

func TestStoreBinary(t *testing.T) {


	for _, c := range Cases {
		t.Run(c.Name, func(t *testing.T) {
			err := StoreBinary(c.Store, c.FilePath)
			if err != nil {
				t.Errorf("%s: unexpected error: %v", c.Name, err)
			}

			// Verify file is created
			_, err = os.Stat(c.FilePath)
			if os.IsNotExist(err) {
				t.Errorf("%s: expected file to be created, but it was not", c.Name)
			}

			// Verify stored data
			file, err := os.Open(c.FilePath)
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

			if len(loadedStore) != len(c.Store) {
				t.Errorf("%s: expected map size %d, got %d", c.Name, len(c.Store), len(loadedStore))
			}
		})
	}
}
