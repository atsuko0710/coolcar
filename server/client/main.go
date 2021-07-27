package main

import (
	"context"
	"fmt"
	"log"

	trippb "coolcar/server/proto/gen/go"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}

	tsClient := trippb.NewTripServiceClient(conn)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip123",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip:%v", err)
	}

	fmt.Println(r)
}
