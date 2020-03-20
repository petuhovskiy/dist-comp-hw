package auth

import "time"

// ValidateRequest is used for access_token validation.
type ValidateRequest struct {
	AccessToken string `json:"access_token"`
}

// ValidateResponse contains details about valid access token.
type ValidateResponse struct {
	UserID      uint          `json:"user_id"`
	ExpireAfter time.Duration `json:"expire_after"`
}


func (c *Client) Validate(token string) (ValidateResponse, error) {
	req := ValidateRequest{
		AccessToken: token,
	}
	var resp ValidateResponse
	err := c.doPost("/v1/validate", &req, &resp)
	return resp, err
}