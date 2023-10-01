package main

import (
	pb "github.com/maksimUlitin/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":8080"

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start %v", err)
	}
}
