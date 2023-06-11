package response

type BioKYC struct {
	response

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
