package main

import (
	"context"
	"log"
	"net"
	"net/http"

	trippb "coolcar/server/proto/gen/go"
	trip "coolcar/server/tripservice"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}

func startGRPCGateway() {
	c := context.Background()
	c, cannel := context.WithCancel(c)
	defer cannel()

	mux := runtime.NewServeMux()
	// 与 8081 建立链接
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cannot start grpc gateway:%v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server:%v", err)
	}
}
 