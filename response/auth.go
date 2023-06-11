package response

type Auth struct {
	response

	// Access token to be used for further APIs calls
	AccessToken string `json:"access_token"`
}
