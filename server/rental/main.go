package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	rentalpb "coolcar/server/rental/api/gen/v1"
	"coolcar/server/rental/trip"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})
	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}
