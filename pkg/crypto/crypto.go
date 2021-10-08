package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Encode a sequence of bytes into a string.
func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Decode a string into a sequence of bytes.
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func cipherFeedbackMode(key string, text []byte) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(text))
	cfb.XORKeyStream(cipherText, text)
	return cipherText
}

// Encrypt a text string using Cipher Feedback Mode.
func Encrypt(key, text string) string {
	plainText := []byte(text)
	cipherText := cipherFeedbackMode(key, plainText)
	return encodeBase64(cipherText)
}

// Decrypt a text string using Cipher Feedback Mode.
func Decrypt(key, text string) string {
	cipherText := decodeBase64(text)
	plainText := cipherFeedbackMode(key, cipherText)
	return string(plainText)
}
