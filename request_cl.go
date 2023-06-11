package nishan

import "github.com/IamFaizanKhalid/nishan-go/errors"

type MobileRequest struct {
	// 13 digit citizen number
	CitizenNo string `json:"citizenNo"`

	// Base64 encoded string of fingerprint
	Fingerprint string `json:"fingerprint"`

	// Finger index from 1 to 10
	FingerIndex int `json:"fingerIndex"`

	// A pakistani mobile number
	MobileNumber string `json:"mobileNumber"`

	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`

	// IMEI of device
	IMEI string `json:"imei"`
}

func (r *MobileRequest) Validate() errors.ErrorCode {
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

	// mobile
	if !isMobileNumberValid(r.MobileNumber) {
		return errors.MobileNumInvalid
	}

	// lat-long
	if !isLatitudeValid(r.Latitude) {
		return errors.LatitudeInvalid
	}
	if !isLongitudeValid(r.Longitude) {
		return errors.LongitudeInvalid
	}

	// imei
	if !isImeiValid(r.IMEI) {
		return errors.ImeiInvalid
	}

	return errors.Nil
}
