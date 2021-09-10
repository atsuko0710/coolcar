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
	CarManager     CarManager
	ProfileManager ProfileManager
	Logger         *zap.Logger
	Mongo          *dao.Mongo
}

// 防入侵
type ProfileManager interface {
	Verify(context.Context, id.AccountID) (id.IdentityID, error)
}

type CarManager interface {
	Verify(context.Context, id.CarID, *rentalpb.Location) error
	Unlock(context.Context, id.CarID) error
}

func (s *Service) CreateTrip(ctx context.Context, in *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 验证驾驶证身份
	iID, err := s.ProfileManager.Verify(ctx, aid)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	// 检查车辆状态
	carId := id.CarID(in.CarId)
	err = s.CarManager.Verify(ctx, carId, in.Start.Location)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	// 先创建行程，避免车开走了但没有记录
	
	
	s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId: aid.String(),
		CarId: carId.String(),
		IdentityId: iID.String(),
		Status: rentalpb.TripStatus_IN_PROGRESS,
	})

	// 车辆开锁

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

	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(c, tid, aid)
	if req.Current != nil {
		// 修改 Current 的位置
		tr.Trip.Current = s.calcCurrentStataus(tr.Trip, req.Current.Location)
	}
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	s.Mongo.UpdateTrip(c, tid, aid, tr.UpdateAt, tr.Trip)
	return nil, status.Error(codes.Unimplemented, "")
}

// calcCurrentStataus 计算当前状态
func (s *Service) calcCurrentStataus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
