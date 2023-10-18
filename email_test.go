package kekasigohelper

import "testing"

func TestIsEmailValid(t *testing.T) {
	if err := IsEmailValid("info@kekasi.co.id"); err != nil {
		t.Error(err)
	}
}
