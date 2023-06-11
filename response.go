package nishan

import "github.com/IamFaizanKhalid/nishan-go/errors"

type common struct {
	// Response status code
	ErrCode errors.ErrorCode `json:"code"`

	// Response status message
	ErrMessage string `json:"message"`
}

type Auth struct {
	common

	// Access token to be used for further APIs calls
	AccessToken string `json:"access_token"`
}

type Bioverisys struct {
	common

	// Array of available finger indexes
	FingerList []int `json:"fingerList"`

	// Transaction id
	TransactionId string `json:"transactionId"`
}

type BioKYC struct {
	common

	// Array of available finger indexes
	FingerList []int `json:"fingerList"`

	// Transaction id
	TransactionId string `json:"transactionId"`

	// Citizen's personal info
	PersonalData PersonalData `json:"personalData"`
}

type PersonalData struct {
	Name       string `json:"name"`
	FatherName string `json:"father_name"`
}
