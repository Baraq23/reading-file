package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

type Entry struct {
	Value  []byte
	Offset int64
}

// StoreBinary serializes the indexed data into a .idx file at the specified filePath.
func StoreBinary(indexed map[uint64]Entry, filePath string) error {

	if filePath == "" {
		return fmt.Errorf("Filepath not provided.")
	}

	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(indexed)
}
