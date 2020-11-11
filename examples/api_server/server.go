package main

import (
	"context"
	"fmt"
	v1 "github.com/clutonik/elastic-cloud-tools/pkg/api/v1"
	"log"
	"net"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) Ping(ctx context.Context, request *v1.PingRequest) (*v1.PingResponse, error){
	address := request.GetConfig().GetClusterAddress()
	// TODO: Use elastic API to ping cluster
	fmt.Printf("Ping %v\n", address)
	result := &v1.PingResponse{Result: "ok"}
	return result, nil
}

func main(){
	fmt.Println("Example code for GRPC server")

	// Start listener
	lis, err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to start GRPC server: %v", err)
	}
	fmt.Printf("Listening on: %v \n", lis.Addr())

	s := grpc.NewServer()
	v1.RegisterClusterServiceServer(s, &server{})

	// Serve the registered services
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve the registered services: %v", err)
	}
}


