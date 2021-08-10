package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	rentalpb "coolcar/server/rental/api/gen/v1"
	"coolcar/server/rental/trip"
	"coolcar/server/shared/auth"
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

	// 生成拦截器
	in, err := auth.Interceptor("shared/auth/public.key")
	if err != nil {
		log.Fatalf("cannot create auth interceptor: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})
	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}
