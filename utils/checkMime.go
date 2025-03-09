package utils

import "net/http"

func CheckMime(file []byte) bool {
	allowedMIMEs := map[string]struct{}{
		"text/plain":                {},
		"text/plain; charset=utf-8": {},
		"application/pdf":           {},
	}

	mimeType := http.DetectContentType(file)
	_, valid := allowedMIMEs[mimeType]

	return valid
}
