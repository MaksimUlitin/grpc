package main

import (
	"context"
	"log"
	"time"

	pb "github.com/maksimUlitin/proto"
)

func callSeyHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SeyHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("cloud not green %v", err)
	}
	log.Printf("%s", res.Message)
}
