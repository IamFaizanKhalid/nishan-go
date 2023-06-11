package response

type Bioverisys struct {
	response

	// Array of available finger indexes
	FingerList []int `json:"fingerList"`

	// Transaction id
	TransactionId string `json:"transactionId"`
}
