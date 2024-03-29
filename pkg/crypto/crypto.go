package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

// EncryptionKey tries to get the value from the PASSAGER_ENCRYPTION_KEY environment variable.
// When the environment variable is not defined it returns a default value.
func EncryptionKey() string {
	if key := os.Getenv("PASSAGER_ENCRYPTION_KEY"); key != "" {
		return key
	}
	return "#encryption-key#"
}

func newGCM(key string) cipher.AEAD {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	return aesgcm
}

// Encrypt a text string using AES.
func Encrypt(text, key string) string {
	gcm := newGCM(key)
	nonce := make([]byte, gcm.NonceSize())
	return string(gcm.Seal(nonce, nonce, []byte(text), nil))
}

// Decrypt a text string using AES.
func Decrypt(text, key string) string {
	gcm := newGCM(key)
	data := []byte(text)
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		log.Fatal(err)
	}
	return string(plainText)
}
