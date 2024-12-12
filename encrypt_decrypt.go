package kekasigohelper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

// PadKey memastikan kunci sesuai dengan ukuran AES
func PadKey(secretKey string) []byte {
	key := []byte(secretKey)
	if len(key) < 16 {
		key = append(key, make([]byte, 16-len(key))...) // Padding untuk 16-byte
	} else if len(key) < 24 {
		key = append(key, make([]byte, 24-len(key))...) // Padding untuk 24-byte
	} else if len(key) < 32 {
		key = append(key, make([]byte, 32-len(key))...) // Padding untuk 32-byte
	}
	return key[:32] // Potong hingga 32-byte jika terlalu panjang
}

// Encrypt menggunakan AES-GCM dengan secret key
func Encrypt(plainText, secretKey string) (string, error) {
	key := PadKey(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt menggunakan AES-GCM dengan secret key
func Decrypt(cipherText, secretKey string) (string, error) {
	key := PadKey(secretKey)

	// Remove any invalid characters (if known issues like spaces or newlines)
	base64Str := strings.TrimSpace(cipherText)
	// Add padding if necessary
	if len(base64Str)%4 != 0 {
		base64Str += strings.Repeat("=", 4-(len(base64Str)%4))
	}

	data, err := base64.URLEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("cipherText terlalu pendek")
	}

	nonce, cipherTextBytes := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
