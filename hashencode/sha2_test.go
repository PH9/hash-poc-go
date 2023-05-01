package hashencode_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"wasith.me/hasher/hashencode"
)

func TestSHA256ASCII85(t *testing.T) {
	r := hashencode.SHA256ASCII85("1111111111111")

	assert.Equal(t, "o<QX[8:/F57o\"c$J#'a\":Hp>7l0hOY1%^,eIi`/J", r)
}

func TestSHA256_192ASCII85(t *testing.T) {
	r := hashencode.SHA256_192ASCII85("1111111111111")

	assert.Equal(t, "o<QX[8:/F57o\"c$J#'a\":Hp>7l0hOY", r)
}

func TestSHA512_240ASCII85(t *testing.T) {
	r := hashencode.SHA512_240ASCII85("1111111111111")

	assert.Equal(t, "\\*D5oRD504,CJ`fKGFlso@7mbBP>hJaT2D#?NpK1", r)
}

func TestSHA512Base64(t *testing.T) {
	r := hashencode.SHA512_240Base64("1111111111111")

	assert.Equal(t, "t+rN3Zm/mcojfOlmhBYJrvPVcwRoaNKFyQEA4V8G", r)
}
