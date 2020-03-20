package service

import (
	"auth/config"
	"auth/modelapi"
	"auth/modeldb"
	"auth/repos/psql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	bcryptCost = 14
)

type Auth struct {
	refreshGen  *Generator
	refreshRepo *psql.RefreshTokens
	usersRepo   *psql.Users
	jwt         *JWT
	conf        *config.Config
}

func NewAuth(refreshGen *Generator, refreshRepo *psql.RefreshTokens, usersRepo *psql.Users, jwt *JWT, conf *config.Config) *Auth {
	return &Auth{
		refreshGen:  refreshGen,
		refreshRepo: refreshRepo,
		usersRepo:   usersRepo,
		jwt:         jwt,
		conf:        conf,
	}
}

func (s *Auth) Signin(req modelapi.SigninRequest) (modelapi.IssuedTokens, error) {
	user, err := s.usersRepo.FindByEmail(req.Email)
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
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return modelapi.SignupResponse{}, err
	}

	user, err := s.usersRepo.Create(modeldb.User{
		Email:        req.Email,
		PasswordHash: hash,
	})
	if err != nil {
		return modelapi.SignupResponse{}, err
	}

	return modelapi.SignupResponse{
		UserID: user.ID,
		Email:  user.Email,
	}, nil
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
		UserID:    user.ID,
		Token:     refreshToken,
		ExpireAt:  expiresAt,
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