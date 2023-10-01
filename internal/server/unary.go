package main

import (
	"context"
	pb "github.com/maksimUlitin/proto"
)

func (s *helloServer) SeyHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil

}
