package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alessioalex/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// initialise a grpc connection
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create our client
	client := proto.NewStreamingServiceClient(conn)

	// initialise the client stream
	ctx := context.Background()
	stream, err := client.LogStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// send some log messages
	for i := range 5 {
		req := proto.LogStreamRequest{
			Timestamp: timestamppb.New(time.Now()),
			Level:     proto.LogLevel_LOG_LEVEL_INFO,
			Message:   fmt.Sprintf("Hello log: %d", i),
		}

		if err := stream.Send(&req); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}

	// close the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	// log the response from the server
	log.Printf("number of sent: %d", res.GetEntriesLogged())
}
