package grpcserv

import (
	"auth/modelapi"
	"auth/service"
	"context"
	"github.com/golang/protobuf/ptypes"
	"lib/pb"
)

type Auth struct {
	auth *service.Auth
}

func NewAuth(auth *service.Auth) *Auth {
	return &Auth{
		auth: auth,
	}
}

func (a *Auth) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	res, err := a.auth.Validate(modelapi.ValidateRequest{
		AccessToken: req.GetAccessToken(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.ValidateResponse{
		UserId:      uint64(res.UserID),
		ExpireAfter: ptypes.DurationProto(res.ExpireAfter),
		Role:        res.Role,
	}, nil
}
