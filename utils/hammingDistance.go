package utils

import "math/bits"

func HammingDistance(a, b uint64) int {
	return bits.OnesCount64(a ^ b)
}