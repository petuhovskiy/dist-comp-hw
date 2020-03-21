package modelapi

import "time"

// RefreshToken is used for issuing new refresh and access tokens.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// IssuedTokens contains just issued refresh and access tokens.
type IssuedTokens struct {
	AccessToken    string `json:"access_token"`
	AccessTokenTTL time.Duration `json:"access_token_ttl"`

	RefreshToken    string `json:"refresh_token"`
	RefreshTokenTLL time.Duration `json:"refresh_token_ttl"`
}
