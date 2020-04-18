package modelapi

import (
	"lib/pb"
	"time"
)

// ValidateRequest is used for access_token validation.
type ValidateRequest struct {
	AccessToken string `json:"access_token"`
}

// ValidateResponse contains details about valid access token.
type ValidateResponse struct {
	UserID      uint          `json:"user_id"`
	ExpireAfter time.Duration `json:"expire_after"`
	Role        pb.AuthRole
}
