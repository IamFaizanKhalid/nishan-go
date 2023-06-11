package nishan

import (
	"github.com/IamFaizanKhalid/nishan-go/errors"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"
)

// UpdateAccessToken will renew access token that will be used to access further APIs
func (c *nishan) UpdateAccessToken() (resp response.Auth) {
	err := c.request(http.MethodGet, "/authenticate", nil, resp)
	if err != nil {
		resp.Error = errors.SendingRequest
		return
	}

	c.accessToken = resp.AccessToken

	return
}
