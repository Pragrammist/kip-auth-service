package utils

import "strings"

func GetFirstPartOfEmail(email string) string {
	at := strings.LastIndex(email, "@")
	if at >= 0 {
		return email[:at]
	}

	return email
}
