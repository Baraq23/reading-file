package utils

import (
	"errors"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("error reading content of the file")
	}
	return file, nil
}
