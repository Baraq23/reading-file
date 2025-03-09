package utils

func SimHash(features []string) uint64 {
	const hashBits = 64 // Standard 64-bit hash
	var weights [hashBits]int

	for _, feature := range features {
		featureHash := HashToken(feature)

		// Update weights based on feature hash
		for i := 0; i < hashBits; i++ {
			if (featureHash & (1 << i)) != 0 {
				weights[i]++
			} else {
				weights[i]--
			}
		}
	}

	// Generate final SimHash
	var simHash uint64
	for i := 0; i < hashBits; i++ {
		if weights[i] > 0 {
			simHash |= (1 << i)
		}
	}

	return simHash
}
