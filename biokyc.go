package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"github.com/IamFaizanKhalid/nishan-go/request"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"
)

func (c *nishan) BioKYC(req request.Bioverisys) (resp response.BioKYC) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisyskyc/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	return
}

func (c *nishan) ContactlessBioKYC(req request.ContactlessBioverisys) (resp response.BioKYC) {
	resp.ErrCode = req.Validate()
	if resp.ErrCode != errors.Nil {
		return
	}

	err := c.request(http.MethodPost, "/bioverisysclkyc/verify", req, &resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	return
}
