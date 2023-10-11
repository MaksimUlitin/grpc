package main

import (
	"log"
	"time"

	pb "github.com/maksimUlitin/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("request with names: %v ", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.send; err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
