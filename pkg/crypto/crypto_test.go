package crypto

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "passwordmanager1"
	text := "The quick brown fox jumps over the lazy dog"
	encryptedText := Encrypt(text, key)
	decryptedText := Decrypt(encryptedText, key)
	assert.Equal(t, text, decryptedText, "texts should be equal")
}

func TestEncryptionKeyNoEnvvar(t *testing.T) {
	os.Setenv("PASSAGER_ENCRYPTION_KEY", "testing-foo-bar!")
	assert.Equal(t, "testing-foo-bar!", EncryptionKey())
}

func TestEncryptionKeyDefaultValue(t *testing.T) {
	os.Unsetenv("PASSAGER_ENCRYPTION_KEY")
	assert.Equal(t, "#encryption-key#", EncryptionKey())
}
