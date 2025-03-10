package utils

func Tokenizer(data []byte) []string {
	var tokens []string
	var currentToken []byte

	for _, b := range data {
		if b == ' ' { // Changed to single quotes for byte literal.
			if len(currentToken) > 0 {
				tokens = append(tokens, string(currentToken))
				currentToken = nil
			}
		} else {
			currentToken = append(currentToken, b)
		}
	}

	if len(currentToken) > 0 {
		tokens = append(tokens, string(currentToken))
	}

	return tokens

}
