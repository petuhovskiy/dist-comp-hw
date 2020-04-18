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

// @Summary Confirms user account phone or email.
// @Tags auth
// @Produce  json
// @Param v query string true "Confirmation string"
// @Success 200 {object} modelapi.ConfirmResponse
// @Router /v1/confirm [get]
func (h *Auth) Confirm(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	resp, err := h.auth.Confirm(v)
	if err != nil {
		log.Println("failed to confirm, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}

// @Summary Updates role of the user
// @Tags management
// @Accept  json
// @Produce  json
// @Param req body modelapi.SetRole true "Request"
// @Success 200 {object} modelapi.SetRole
// @Router /v1/setrole [post]
// @Security ApiKeyAuth
func (h *Auth) SetRole(w http.ResponseWriter, r *http.Request) {
	var data modelapi.SetRole
	if err := render.Decode(r, &data); err != nil {
		log.Println("failed to read request, err=", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	resp, err := h.auth.SetRole(data)
	if err != nil {
		log.Println("failed to refresh, err=", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Respond(w, r, resp)
}
