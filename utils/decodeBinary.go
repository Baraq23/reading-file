package utils

import (
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"
)

func DecodeBinary(path string) (map[uint64]Entry, error) {
	if filepath.Ext(path) != ".idx" {
		return nil, errors.New("wrong type of file")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error while decoding file")
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var result map[uint64]Entry
	if err := decoder.Decode(&result); err != nil {
		return nil, errors.New("error while decoding file")
	}

	return result, nil
}
