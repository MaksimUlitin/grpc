package main

import (
	"context"
	pb "github.com/maksimUlitin/proto"
	"io"
	"log"
)

func CallSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming starting")
	stream, err := client.SeyHelloClientStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names %v", names)
	}
	for {
		message, err := sream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("sreaming finished")
}
