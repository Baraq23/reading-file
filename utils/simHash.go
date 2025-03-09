package utils

import (
	"crypto/md5"
)

func SimHash(data []byte) (simhash uint64 ){
	hash := md5.Sum(data) 
	const n = 8

	for i := 0; i < n; i++ {
		simhash |= uint64(hash[i]) << (i * n)
	}
	
	return 
}
