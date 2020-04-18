package service

import (
	"auth/modeldb"
	"encoding/json"
	"fmt"
	"github.com/cristalhq/jwt"
	"lib/pb"
	"time"
)

const (
	jwtIssuer = "auth-server"
)

type Claims struct {
	jwt.StandardClaims
	Role pb.AuthRole
}

type JWT struct {
	signer  jwt.Signer
	builder *jwt.TokenBuilder
	gen     *Generator
}

func NewJWT(signer jwt.Signer, builder *jwt.TokenBuilder, gen *Generator) *JWT {
	return &JWT{
		signer:  signer,
		builder: builder,
		gen:     gen,
	}
}

func (j *JWT) IssueToken(user modeldb.User, ttl time.Duration) (string, error) {
	id, err := j.gen.Generate()
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(ttl)

	token, err := j.builder.Build(&Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  []string{"user"},
			ExpiresAt: jwt.Timestamp(expiresAt.Unix()),
			ID:        id,
			IssuedAt:  jwt.Timestamp(now.Unix()),
			Issuer:    jwtIssuer,
			Subject:   fmt.Sprint(user.ID),
		},
		Role: user.Role,
	})
	if err != nil {
		return "", err
	}

	return token.InsecureString(), nil
}

func (j *JWT) ValidateToken(raw string) (*Claims, error) {
	token, err := jwt.ParseAndVerifyString(raw, j.signer)
	if err != nil {
		return nil, err
	}

	claims := &Claims{}
	err = json.Unmarshal(token.RawClaims(), claims)
	if err != nil {
		return nil, err
	}

	validator := jwt.NewValidator(
		jwt.ExpirationTimeChecker(time.Now()),
	)
	err = validator.Validate(&claims.StandardClaims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
