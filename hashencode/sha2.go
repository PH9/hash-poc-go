package hashencode

import (
	"encoding/ascii85"
	"encoding/base64"

	"wasith.me/hasher/hasher"
)

// SHA256ASCII85
func SHA256ASCII85(original string) string {
	bs := hasher.SHA256(original)
	buffer := make([]byte, ascii85.MaxEncodedLen(len(bs)))
	_ = ascii85.Encode(buffer, bs)
	return string(buffer)
}

// SHA256_192ASCII85
func SHA256_192ASCII85(original string) string {
	bs := hasher.SHA256_192(original)
	buffer := make([]byte, ascii85.MaxEncodedLen(len(bs)))
	_ = ascii85.Encode(buffer, bs)
	return string(buffer)
}

func SHA512_240ASCII85(original string) string {
	bs := hasher.SHA512_240(original)
	buffer := make([]byte, ascii85.MaxEncodedLen(len(bs)))
	_ = ascii85.Encode(buffer, bs)
	return string(buffer)
}

func SHA512_240Base64(origin string) string {
	bs := hasher.SHA512_240(origin)
	return base64.StdEncoding.EncodeToString(bs)
}
