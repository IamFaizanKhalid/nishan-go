package request

import (
	"encoding/base64"
	"regexp"
)

func isCnicValid(s string) bool {
	return len(s) == 13 && isNumeric(s)
}

func isImeiValid(s string) bool {
	return len(s) == 15 && isNumeric(s)
}

var mobileRegex = regexp.MustCompile(`^(((\+?92)|0) ?)?3[0-9]{2} ?[0-9]{7}$`)

func isMobileNumberValid(s string) bool {
	return mobileRegex.MatchString(s)
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

func isLatitudeValid(lat float64) bool {
	return lat >= -90 && lat <= 90
}

func isLongitudeValid(lat float64) bool {
	return lat >= -180 && lat <= 180
}
