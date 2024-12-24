package kekasigohelper

import "testing"

func TestHashPassword(t *testing.T) {
	plainText := "Arditya Kekasi"

	HashBcrypt(&plainText)
	LoggerInfo(plainText)
}
