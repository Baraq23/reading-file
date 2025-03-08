package utils

// Chunk splits a byte slice into smaller chunks of the specified size.
func Chunk(content []byte, sep int) [][]byte {

	var chunks [][]byte

	for i := 0; i < len(content); i += sep {
		end := i + sep
		if end > len(content) {
			end = len(content)
		}
		chunks = append(chunks, content[i:end])
	}

	return chunks
}