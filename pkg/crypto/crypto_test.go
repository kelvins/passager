package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "a1a2a3a4a5a6a7a8"
	text := "The quick brown fox jumps over the lazy dog"
	encryptedText := Encrypt(key, text)
	decryptedText := Decrypt(key, encryptedText)
	assert.Equal(t, text, decryptedText, "texts should be equal")
}
