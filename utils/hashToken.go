package utils

import "hash/fnv"

func HashToken(str string) uint64 {
	return fnv.New64a().Sum64()
}
