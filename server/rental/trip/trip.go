package trip

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rentalpb "coolcar/server/rental/api/gen/v1"
	"coolcar/server/rental/trip/dao"

	"coolcar/server/shared/auth"
	"coolcar/server/shared/id"
)

type Service struct {
	Logger *zap.Logger
	Mongo  *dao.Mongo
}

func (s *Service) CreateTrip(ctx context.Context, in *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	// aid, err := auth.AccountIDFromContext(ctx)
	_, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if req.Current != nil {
		
	}
	return nil, status.Error(codes.Unimplemented, "")
}
