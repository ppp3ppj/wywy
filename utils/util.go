package utils

import (
	"net/mail"
)

func ValidateEmailFormat(email string) (string, bool) {
    addr, err := mail.ParseAddress(email)
    if err != nil {
        return "", false
    }
    return addr.Address, true
}


