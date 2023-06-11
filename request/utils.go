package request

import "encoding/base64"

func isCnicValid(s string) bool {
	return len(s) == 13 && isNumeric(s)
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || '9' < c {
			return false
		}
	}
	return true
}

func isBase64Encoded(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
