package auth

// type AuthServiceServer interface {
// 	Login(context.Context, *LoginRequest) (*LoginResponse, error)
// }
import (
	"context"
	authpb "coolcar/server/auth/api/gen/v1"
	"coolcar/server/auth/dao"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	OpenIDResolver OpenIDResolver
}

// OpenIDResolver resolvers an authorization code to an open id
type OpenIDResolver interface {
	Resolver(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// s.Logger.Info("receive code:", zap.String("code", req.Code))

	openID, err := s.OpenIDResolver.Resolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolver open id:%v", err)
	}

	accountId, err := s.Mongo.ResolveAccountId(c, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id:", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "")
	}

	tkn, err := s.TokenGenerator.GenerateToken(accountId, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token:", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
