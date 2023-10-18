package kekasigohelper

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

func IsEmailValid(email string) error {
	// Check email format
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	// Split email to get domain
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return errors.New("Domain not found")
	}

	// Check domain MX records
	_, err = net.LookupMX(parts[1])
	if err != nil {
		return err
	}

	return nil
}
