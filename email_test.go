package kekasigohelper

import "testing"

func TestIsEmailValid(t *testing.T) {
	if err := IsEmailValid("info@kekasi.co.id"); err != nil {
		t.Error(err)
	}
}

func TestIsEmailDomainEqual(t *testing.T) {
	if IsEmailDomainEqual("info@kekasi.co.id", "arditya@kekasi.co.id") {
		LoggerInfo("Email domain is equal")
	}
}
