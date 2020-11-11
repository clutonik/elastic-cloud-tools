package main

import (
	"context"
	v1 "github.com/clutonik/elastic-cloud-tools/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
)

func main(){
	// Connect to GRPC server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to GRPC server: &v", err)
	}
	defer cc.Close() // Close client connection

	// Create a client now
	c := v1.NewClusterServiceClient(cc)

	// Invoke Ping service on GRPC Server
	invokePing(c)

}

func invokePing(client v1.ClusterServiceClient){
	// Use Ping Service
	req := &v1.PingRequest{
		Config:
		&v1.Cluster_Config{
			ClusterAddress:     "localhost",
			UserName:           "test",
			Password:           "test",
			DeploymentTemplate: 0,
		},
	}
	res, err := client.Ping(context.Background(), req)
	if err != nil {
		log.Fatalf("could not ping cluster: %v", err)
	}
	log.Printf("response: %v", res.GetResult())
}