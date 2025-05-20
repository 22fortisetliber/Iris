package util

import (
	"crypto/sha256"
)

// Hash compute SHA1 hashes of s given input.
func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
