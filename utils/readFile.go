package utils

import (
	"bytes"
	"errors"
	"os"
	"unicode/utf8"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("error reading content of the file")
	}

	utf8BOM := []byte{0xEF, 0xBB, 0xBF}
	file = bytes.TrimPrefix(file, utf8BOM)

	if !utf8.Valid(file) {
		return nil, errors.New("file contains invalid UTF-8")
	}

	return file, nil
}
