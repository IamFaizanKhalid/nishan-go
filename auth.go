package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"
)

func (c *nishan) UpdateAccessToken() (resp response.Auth) {
	err := c.request(http.MethodGet, "/authenticate", nil, resp)
	if err != nil {
		resp.ErrCode = errors.SendingRequest
		return
	}

	c.accessToken = resp.AccessToken

	return
}
