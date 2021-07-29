package main

import (
	"log"
	"net"

	authpb "coolcar/server/auth/api/gen/v1"
	"coolcar/server/auth/auth"
	"coolcar/server/wechat"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		OpenIDResolver: &wechat.Service{
			AppID: "wx2686c8410b229d6f",
			AppSecret: "db2905a97cc3409549141ae2446966e5",
		},
		Logger: logger,
	})
	err = s.Serve(lis)
	logger.Fatal("cannot server", zap.Error(err))
}
