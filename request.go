package nishan

import "github.com/IamFaizanKhalid/nishan-go/errors"

type Request struct {
	// 13 digit citizen number
	CitizenNo string `json:"citizenNo"`

	// Base64 encoded string of fingerprint
	Fingerprint string `json:"fingerprint"`

	// Finger index from 1 to 10
	FingerIndex int `json:"fingerIndex"`
}

func (r *Request) Validate() errors.ErrorCode {
	// citizen number
	if r.CitizenNo == "" {
		return errors.CitizenNumEmpty
	}
	for !isCnicValid(r.CitizenNo) {
		return errors.CitizenNumInvalid
	}

	// fingerprint
	if r.Fingerprint == "" {
		return errors.FingerprintEmpty
	}
	if !isBase64Encoded(r.Fingerprint) {
		return errors.FingerprintFormat
	}

	if r.FingerIndex < 1 || r.FingerIndex > 10 {
		return errors.FingerIndexInvalid
	}

	return errors.Nil
}
