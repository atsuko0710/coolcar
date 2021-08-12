package trip

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rentalpb "coolcar/server/rental/api/gen/v1"

	"coolcar/server/shared/auth"
)

type Service struct {
	Logger         *zap.Logger
}

func (s *Service) CreateTrip(ctx context.Context, in *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	// aid, err := auth.AccountIDFromContext(ctx)
	_, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return nil, status.Error(codes.Unimplemented, "")
}
