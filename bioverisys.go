package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"github.com/IamFaizanKhalid/nishan-go/request"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"
)

// Bioverisys API can be used to verify fingerprint.
func (c *nishan) Bioverisys(req request.BioverisysRequest) (resp response.Bioverisys) {
	resp.Error = req.Validate()
	if resp.Error != errors.None {
		return
	}

	err := c.request(http.MethodPost, "/bioverisys/verify", req, &resp)
	if err != nil {
		resp.Error = errors.SendingRequest
		return
	}

	return
}
