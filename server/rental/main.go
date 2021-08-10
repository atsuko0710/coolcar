package main

import (
	"log"

	rentalpb "coolcar/server/rental/api/gen/v1"
	"coolcar/server/rental/trip"
	"coolcar/server/shared/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Logger:             logger,
		Addr:               ":8082",
		Name:               "rental",
		AuthPublicKeyFiled: "shared/auth/public.key",
		RegisterServiceFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
	}))
}
