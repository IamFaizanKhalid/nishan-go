package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"github.com/IamFaizanKhalid/nishan-go/request"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"
)

// Bioverisys API can be used to verify fingerprint
func (c *nishan) Bioverisys(req request.Bioverisys) (resp response.Bioverisys) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisys/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	return
}

// ContactlessBioverisys API can be used to verify fingerprint
func (c *nishan) ContactlessBioverisys(req request.ContactlessBioverisys) (resp response.Bioverisys) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisyscl/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	return
}
