package response

import "github.com/IamFaizanKhalid/nishan-go/errors"

type response struct {
	// Response status code
	Error errors.Error `json:"code"`

	// Response status message
	Message string `json:"message"`
}
