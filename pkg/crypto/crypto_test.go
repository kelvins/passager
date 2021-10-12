package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "passwordmanager1"
	text := "The quick brown fox jumps over the lazy dog"
	encryptedText := Encrypt(text, key)
	decryptedText := Decrypt(encryptedText, key)
	assert.Equal(t, text, decryptedText, "texts should be equal")
}
