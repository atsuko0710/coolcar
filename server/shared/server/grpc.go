package server

import (
	"net"

	"coolcar/server/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Logger              *zap.Logger
	Addr                string
	Name                string
	AuthPublicKeyFiled  string
	RegisterServiceFunc func(*grpc.Server)
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption
	if c.AuthPublicKeyFiled != "" {
		// 生成拦截器
		in, err := auth.Interceptor(c.AuthPublicKeyFiled)
		if err != nil {
			c.Logger.Fatal("cannot create auth interceptor", nameField, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)

	c.RegisterServiceFunc(s)

	c.Logger.Info("server started at", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
