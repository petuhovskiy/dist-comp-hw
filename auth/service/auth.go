package service

import (
	"auth/config"
	"auth/modelapi"
	"auth/modeldb"
	"auth/repos/psql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lib/pb"
	"time"
)

const (
	bcryptCost = 14
)

var confirmTTL = time.Hour

type Auth struct {
	refreshGen  *Generator
	refreshRepo *psql.RefreshTokens
	usersRepo   *psql.Users
	confirmGen  *Generator
	confirmRepo *psql.Confirm
	notificator *Notificator
	jwt         *JWT
	conf        *config.Config
}

func NewAuth(refreshGen *Generator, refreshRepo *psql.RefreshTokens, usersRepo *psql.Users, confirmGen *Generator, confirmRepo *psql.Confirm, notificator *Notificator, jwt *JWT, conf *config.Config) *Auth {
	return &Auth{
		refreshGen:  refreshGen,
		refreshRepo: refreshRepo,
		usersRepo:   usersRepo,
		confirmGen:  confirmGen,
		confirmRepo: confirmRepo,
		notificator: notificator,
		jwt:         jwt,
		conf:        conf,
	}
}

func (s *Auth) Signin(req modelapi.SigninRequest) (modelapi.IssuedTokens, error) {
	user, err := s.usersRepo.FindByLogin(req.Login)
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(req.Password))
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	return s.issueTokens(user)
}

func (s *Auth) Signup(req modelapi.SignupRequest) (modelapi.SignupResponse, error) {
	if req.Email == "" && req.Phone == "" {
		return modelapi.SignupResponse{}, ErrLoginIsRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return modelapi.SignupResponse{}, err
	}

	user, err := s.usersRepo.Create(modeldb.User{
		PasswordHash: hash,
		Role:         pb.AuthRole_USER,
	})
	if err != nil {
		return modelapi.SignupResponse{}, err
	}

	if req.Email != "" {
		err := s.requireConfirmation(user.ID, modeldb.ConfirmEmail, req.Email)
		if err != nil {
			return modelapi.SignupResponse{}, err
		}
	}

	if req.Phone != "" {
		err := s.requireConfirmation(user.ID, modeldb.ConfirmSms, req.Phone)
		if err != nil {
			return modelapi.SignupResponse{}, err
		}
	}

	return modelapi.SignupResponse{
		UserID: user.ID,
	}, nil
}

func (s *Auth) requireConfirmation(userID uint, t modeldb.ConfirmType, subj string) error {
	link, err := s.confirmGen.Generate()
	if err != nil {
		return err
	}

	obj, err := s.confirmRepo.Create(
		modeldb.Confirm{
			Link:     link,
			UserID:   userID,
			Type:     t,
			Subject:  subj,
			ExpireAt: time.Now().Add(confirmTTL),
		},
	)
	if err != nil {
		return err
	}

	// content contains link to confirm account
	content := fmt.Sprintf(s.conf.ConfirmUrl, obj.Link)

	return s.notificator.Notify(obj.Type, subj, content)
}

func (s *Auth) Refresh(req modelapi.RefreshRequest) (modelapi.IssuedTokens, error) {
	refresh, err := s.refreshRepo.FindNonExpired(req.RefreshToken)
	if err != nil {
		return modelapi.IssuedTokens{}, err // TODO: return not found or expired token
	}

	user, err := s.usersRepo.FindByID(refresh.UserID)
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	err = s.refreshRepo.Delete(refresh.ID)
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	return s.issueTokens(user)
}

func (s *Auth) Validate(req modelapi.ValidateRequest) (modelapi.ValidateResponse, error) {
	token := req.AccessToken

	claims, err := s.jwt.ValidateToken(token)
	if err != nil {
		return modelapi.ValidateResponse{}, err
	}

	var userID uint
	_, err = fmt.Sscanf(claims.Subject, "%d", &userID)
	if err != nil {
		return modelapi.ValidateResponse{}, err
	}

	return modelapi.ValidateResponse{
		UserID:      userID,
		ExpireAfter: time.Until(claims.ExpiresAt.Time()),
		Role:        claims.Role,
	}, nil
}

func (s *Auth) Confirm(link string) (modelapi.ConfirmResponse, error) {
	c, err := s.confirmRepo.Find(link)
	if err != nil {
		return modelapi.ConfirmResponse{}, err
	}

	switch c.Type {
	case modeldb.ConfirmEmail:
		err = s.usersRepo.UpdateEmail(c.UserID, c.Subject)

	case modeldb.ConfirmSms:
		err = s.usersRepo.UpdatePhone(c.UserID, c.Subject)

	default:
		return modelapi.ConfirmResponse{}, ErrUnknownNotifyType
	}

	if err != nil {
		return modelapi.ConfirmResponse{}, err
	}

	err = s.confirmRepo.Delete(link)
	if err != nil {
		return modelapi.ConfirmResponse{}, err
	}

	return modelapi.ConfirmResponse{
		Subject: c.Subject,
		Message: "Successfully confirmed!",
	}, nil
}

func (s *Auth) issueTokens(user modeldb.User) (modelapi.IssuedTokens, error) {
	accessToken, err := s.jwt.IssueToken(user, s.conf.AccessTokenExpiration)
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	refreshToken, err := s.refreshGen.Generate()
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}
	expiresAt := time.Now().Add(s.conf.RefreshTokenExpiration)

	_, err = s.refreshRepo.Create(modeldb.RefreshToken{
		UserID:   user.ID,
		Token:    refreshToken,
		ExpireAt: expiresAt,
	})
	if err != nil {
		return modelapi.IssuedTokens{}, err
	}

	return modelapi.IssuedTokens{
		AccessToken:     accessToken,
		AccessTokenTTL:  s.conf.AccessTokenExpiration,
		RefreshToken:    refreshToken,
		RefreshTokenTLL: s.conf.RefreshTokenExpiration,
	}, nil
}

func (s *Auth) SetRole(req modelapi.SetRole) (modelapi.SetRole, error) {
	err := s.usersRepo.UpdateRole(req.UserID, req.Role)
	if err != nil {
		return req, err
	}

	return req, nil
}
