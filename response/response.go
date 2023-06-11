package response

import "github.com/IamFaizanKhalid/nishan-go/errors"

type response struct {
	// Response status code
	ErrCode errors.ErrorCode `json:"code"`

	// Response status message
	ErrMessage string `json:"message"`
}
