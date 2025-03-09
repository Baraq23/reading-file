package utils


// CalculateSimHash generates a SimHash fingerprint from a list of features
func SimHash(features []string) (simHash uint64) {
	const hashBits = 8
	weights := make([]float64, hashBits)

	for _, feature := range features {
		featureHash := HashToken(feature)

		// Update weights based on the feature hash
		for i := 0; i < hashBits; i++ {
			if (featureHash & (1 << i)) > 0 {
				weights[i]++
			} else {
				weights[i]--
			}
		}
	}

	// Generate the final SimHash based on the weights
	for i := 0; i < hashBits; i++ {
		if weights[i] > 0 {
			simHash |= (1 << i)
		}
	}

	return
}
