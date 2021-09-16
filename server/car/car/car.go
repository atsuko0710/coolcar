package car

import (
	"context"
	carpb "coolcar/server/car/api/gen/v1"
	"coolcar/server/car/dao"
	"coolcar/server/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	Mongo  *dao.Mongo
}

func (s *Service) CreateCar(ctx context.Context, in *carpb.CreateCarRequest) (*carpb.CarEntity, error) {
	cr, err := s.Mongo.CreateCar(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &carpb.CarEntity{
		Id:  cr.ID.Hex(),
		Car: cr.Car,
	}, nil
}

func (s *Service) GetCar(ctx context.Context, in *carpb.GetCarRequest) (*carpb.Car, error) {
	cr, err := s.Mongo.GetCar(ctx, id.CarID(in.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	return cr.Car, nil
}

func (s *Service) GetCars(ctx context.Context, in *carpb.GetCarsRequest) (*carpb.GetCarsResponse, error) {
	cars, err := s.Mongo.GetCars(ctx)
	if err != nil {
		s.Logger.Error("cannot get cars", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	res := &carpb.GetCarsResponse{}
	for _, cr := range cars {
		res.Cars = append(res.Cars, &carpb.CarEntity{
			Id: cr.ID.Hex(),
			Car: cr.Car,
		})
	}
	return res, nil
}

func (s *Service) LockCar(ctx context.Context, in *carpb.LockCarRequest) (*carpb.LockCarResponse, error) {
	return nil, nil
}

func (s *Service) UnlockCar(ctx context.Context, in *carpb.UnlockCarRequest) (*carpb.UnlockCarResponse, error) {
	return nil, nil
}

func (s *Service) UpdateCar(ctx context.Context, in *carpb.UpdateCarRequest) (*carpb.UpdateCarResponse, error) {
	return nil, nil
}
