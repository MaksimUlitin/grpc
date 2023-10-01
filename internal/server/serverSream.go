package main

import (
	pb "github.com/maksimUlitin/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SeyHelloClientStreamingClient) error {
	log.Printf("request with names: %v ", pb.NamesList{})
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
