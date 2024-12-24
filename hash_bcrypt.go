package kekasigohelper

import "golang.org/x/crypto/bcrypt"

func HashBcrypt(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func CompareBcrypt(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
