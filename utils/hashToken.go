package utils

import "hash/fnv"

func HashToken(feature string) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(feature))
	return hasher.Sum64()
}
