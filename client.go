package nishan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IamFaizanKhalid/nishan-go/errors"
)

type nishan struct {
	baseUrl     string
	apiKey      string
	accessToken string
}

func New(subdomain, apiKey string) (*nishan, error) {
	c := &nishan{
		baseUrl: fmt.Sprintf("https://%s.nadra.gov.pk", subdomain),
		apiKey:  apiKey,
	}

	resp := c.UpdateAccessToken()
	if resp.Error != errors.None {
		return nil, fmt.Errorf("auth error: %s", resp.Message)
	}

	return c, nil
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

	request, err := http.NewRequest(method, c.baseUrl+path, body)
	if err != nil {
		return err
	}

	// set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", c.apiKey)
	request.Header.Set("Token", c.accessToken)

	// hit api
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// parse response
	if requestObj != nil {
		return json.NewDecoder(response.Body).Decode(responseObj)
	}

	return nil
}
