package handlers

import (
	"auth/modelapi"
	"auth/service"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

type Auth struct {
	auth *service.Auth
}

func NewAuthV1(p *service.Auth) *Auth {
	return &Auth{
		auth: p,
	}
}

// @Summary Sign up (register)
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body modelapi.SignupRequest true "Credentials"
// @Success 200 {object} modelapi.SignupResponse
// @Router /v1/signup [post]
func (h *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var data modelapi.SignupRequest
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.auth.Signup(data)
	if err != nil {
		log.Println("failed to sign up, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

// @Summary Sign in (login)
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body modelapi.SigninRequest true "Credentials"
// @Success 200 {object} modelapi.IssuedTokens
// @Router /v1/signin [post]
func (h *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	var data modelapi.SigninRequest
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.auth.Signin(data)
	if err != nil {
		log.Println("failed to sign in, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

// @Summary Accepts valid refresh token, returns new refresh and access tokens.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body modelapi.RefreshRequest true "Refresh token"
// @Success 200 {object} modelapi.IssuedTokens
// @Router /v1/refresh [post]
func (h *Auth) Refresh(w http.ResponseWriter, r *http.Request) {
	var data modelapi.RefreshRequest
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.auth.Refresh(data)
	if err != nil {
		log.Println("failed to refresh, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

// @Summary Validates access token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body modelapi.ValidateRequest true "Access token"
// @Success 200 {object} modelapi.ValidateResponse
// @Router /v1/validate [post]
func (h *Auth) Validate(w http.ResponseWriter, r *http.Request) {
	var data modelapi.ValidateRequest
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.auth.Validate(data)
	if err != nil {
		log.Println("failed to validate, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}