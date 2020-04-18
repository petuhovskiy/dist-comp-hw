package service

import (
	"auth/modeldb"
	"github.com/cristalhq/jwt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"lib/pb"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	jwtIDGenerator := NewGenerator(10)

	// initializing services
	jwtSigner, err := jwt.NewHS256([]byte(`abc`))
	assert.Nil(t, err)
	jwtBuilder := jwt.NewTokenBuilder(jwtSigner)
	jwtService := NewJWT(jwtSigner, jwtBuilder, jwtIDGenerator)

	token, err := jwtService.IssueToken(modeldb.User{
		ID:           1,
		CreatedAt:    time.Now(),
		Email:        "petuhovskiy@yandex.ru",
		Phone:        "",
		PasswordHash: nil,
		Role:         pb.AuthRole_ADMIN,
	}, time.Minute)
	assert.Nil(t, err)
	spew.Dump(token)

	claims, err := jwtService.ValidateToken(token)
	assert.Nil(t, err)

	spew.Dump(claims)
	assert.Equal(t, pb.AuthRole_ADMIN, claims.Role)
}