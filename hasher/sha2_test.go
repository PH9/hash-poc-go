package hasher_test

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"wasith.me/hasher/hasher"
)

func TestSHA256(t *testing.T) {
	r := hasher.SHA256("1111111111111")

	assert.Equal(
		t,
		"f3b2ce1a487bc887474e85927fa4bcca4f3f2ab2e9ef530b31f484447f1e6526",
		fmt.Sprintf("%x", r),
	)
}

func TestSHA256_192(t *testing.T) {
	r := hasher.SHA256_192("1111111111111")

	assert.Equal(
		t,
		"f3b2ce1a487bc887474e85927fa4bcca4f3f2ab2e9ef530b",
		fmt.Sprintf("%x", r),
	)
}

func TestSHA512(t *testing.T) {
	r := hasher.SHA512("1111111111111")

	assert.Equal(
		t,
		"b7eacddd99bf99ca237ce966841609aef3d573046868d285c90100e15f0624bc6832a3d98ad3f0e83cfdbd601096806132dae0bdaa0af6112840c0c5ea41e190",
		fmt.Sprintf("%x", r),
	)
}

func TestSHA512_240(t *testing.T) {
	r := hasher.SHA512_240("1111111111111")

	assert.Equal(
		t,
		"b7eacddd99bf99ca237ce966841609aef3d573046868d285c90100e15f06",
		fmt.Sprintf("%x", r),
	)
}

func TestHash(t *testing.T) {
	h := sha256.New()
	originalText := "1111111111111"
	r := hasher.Hash(originalText, h, 24)

	expected := hasher.SHA256_192(originalText)
	assert.Equal(
		t,
		expected,
		r,
	)
}
