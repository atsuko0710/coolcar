package auth

// type AuthServiceServer interface {
// 	Login(context.Context, *LoginRequest) (*LoginResponse, error)
// }
import (
	"context"
	authpb "coolcar/server/auth/api/gen/v1"

	"go.uber.org/zap"
)

type Service struct{
	Logger *zap.Logger
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("receive code:", zap.String("code", req.Code))
	return &authpb.LoginResponse{
		AccessToken: "token receive" + req.Code,
		ExpiresIn: 7200,
	}, nil
}