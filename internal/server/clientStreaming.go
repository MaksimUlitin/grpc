package main

import (
	pb "github.com/maksimUlitin/proto"
	"io"
	"log"
)
var messages []string
func (s *helloServer) seyHelloClientStreaming(stream pb.GreetService_SeyHelloClientStreamingServer) error {

	for {
	req, err := stream.Recv()
	if err == io.EOF{
		return  stream.SendAndClose(&pb.MessagesList{Message: messages})
	} if err != nil {
		return err
		}
		log.Printf("got request with name: %v ", req.Name)
	log.Printf("got request with name: %v" req.Name)
	messages = append(messages, "Hello", req.Name)
	}
}
