package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"net/http"
)

func (c *nishan) VerifyFingerprint(req Request) (resp Bioverisys) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisys/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
	}

	return
}

func (c *nishan) VerifyFingerprintFromMobile(req MobileRequest) (resp Bioverisys) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisyscl/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
	}

	return
}
