package hasher

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

// SHA256 has different nonce with SHA512 but same algorithm (I guess)
func SHA256(original string) []byte {
	h := sha256.New()
	h.Write([]byte(original))
	bs := h.Sum(nil)
	return bs
}

// SHA256_192 is a SHA256 with left most 192bits (I guess)
func SHA256_192(original string) []byte {
	return SHA256(original)[:24]
}

// SHA512 has different nonce with SHA256 but same algorithm (I guess)
func SHA512(original string) []byte {
	h := sha512.New()
	h.Write([]byte(original))
	bs := h.Sum(nil)
	return bs
}

// SHA512_240 is a SHA512 with left most 240bits (I guess)
func SHA512_240(original string) []byte {
	return SHA512(original)[:30]
}

// Hash
func Hash(original string, h hash.Hash, size uint) []byte {
	h.Write([]byte(original))
	return h.Sum(nil)[:size]
}
