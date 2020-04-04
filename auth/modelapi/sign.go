package modelapi

// SigninRequest is used for login.
type SigninRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// SignupRequest is used for registration.
type SignupRequest struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// SignupResponse contains basic user info about just created user.
type SignupResponse struct {
	UserID uint   `json:"user_id"`
}