package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"net/http"
)

func (c *nishan) KnowYourCustomer(req Request) (resp BioKYC) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisyskyc/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
	}

	return
}

func (c *nishan) KnowYourCustomerFromMobile(req MobileRequest) (resp BioKYC) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisysclkyc/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
	}

	return
}
