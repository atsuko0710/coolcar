package trip

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rentalpb "coolcar/server/rental/api/gen/v1"
)

type Service struct {
	Logger         *zap.Logger
}

func (s *Service) CreateTrip(ctx context.Context, in *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	s.Logger.Info("create trip", zap.String("start", in.Start))
	return nil, status.Error(codes.Unimplemented, "")
}
