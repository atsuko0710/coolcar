package auth

// type AuthServiceServer interface {
// 	Login(context.Context, *LoginRequest) (*LoginResponse, error)
// }
import (
	"context"
	authpb "coolcar/server/auth/api/gen/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
}

// OpenIDResolver resolvers an authorization code to an open id
type OpenIDResolver interface {
	Resolver(code string) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// s.Logger.Info("receive code:", zap.String("code", req.Code))

	openID, err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolver open id:%v", err)
	}

	return &authpb.LoginResponse{
		AccessToken: "token receive" + openID,
		ExpiresIn:   7200,
	}, nil
}
