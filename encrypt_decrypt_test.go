package kekasigohelper

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	secretKey := "kekasi.co.id"
	plainText := "Arditya Kekasi"

	// Enkripsi
	encryptedText, err := Encrypt(plainText, secretKey)
	if err != nil {
		t.Error("Error enkripsi : ", err)
	}
	LoggerInfo(encryptedText)
}

func TestDecrypt(t *testing.T) {
	secretKey := "kekasi.co.id"
	encryptText := "DsM0xWEgWkjlGKRRs2Ymac0qd8PLC1bkb3wv5mCLkp3xW6f9zEPVlDU-"

	// Dekripsi
	decryptedText, err := Decrypt(encryptText, secretKey)
	if err != nil {
		t.Error("Error dekripsi : ", err)
	}
	LoggerInfo(decryptedText)
}
