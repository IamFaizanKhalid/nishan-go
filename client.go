package nishan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/IamFaizanKhalid/nishan-go/request"
	"github.com/IamFaizanKhalid/nishan-go/response"
	"net/http"

	"github.com/IamFaizanKhalid/nishan-go/errors"
)

type Client interface {
	UpdateAccessToken() response.Auth
	Bioverisys(request.Bioverisys) response.Bioverisys
	ContactlessBioverisys(request.ContactlessBioverisys) response.Bioverisys

	request(string, string, interface{}, interface{}) error
}

func NewClient(subdomain, apiKey string) (Client, error) {
	c := &nishan{
		baseUrl: fmt.Sprintf("https://%s.nadra.gov.pk", subdomain),
		apiKey:  apiKey,
	}

	resp := c.UpdateAccessToken()
	if resp.ErrCode != errors.Nil {
		return nil, fmt.Errorf("auth error: %s", resp.ErrMessage)
	}

	return c, nil
}

type nishan struct {
	baseUrl     string
	apiKey      string
	accessToken string
}

func (c *nishan) request(method, path string, requestObj interface{}, responseObj interface{}) error {
	// build request
	var body *bytes.Reader
	if requestObj != nil {
		requestBytes, err := json.Marshal(requestObj)
		if err != nil {
			return err
		}
		body = bytes.NewReader(requestBytes)
	}

	req, err := http.NewRequest(method, c.baseUrl+path, body)
	if err != nil {
		return err
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Token", c.accessToken)

	// hit api
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// parse response
	if responseObj != nil {
		return json.NewDecoder(resp.Body).Decode(responseObj)
	}

	return nil
}
