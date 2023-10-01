package main

import (
	"context"
	pb "github.com/maksimUlitin/proto"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming starting")
	stream, err := client.SeyHelloClientStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", names)
	}
	log.Printf("", stream)
}
