package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"net/http"
)

func (c *nishan) UpdateAccessToken() (resp Auth) {
	err := c.request(http.MethodGet, "/authenticate", nil, resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	c.accessToken = resp.AccessToken

	return
}
