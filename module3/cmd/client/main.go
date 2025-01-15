package main

import (
	"context"
	"io"
	"log"

	"github.com/alessioalex/grpc-course/module3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// first initialise grpc connection
	ctx := context.Background()
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// create the client
	client := proto.NewStreamingServiceClient(conn)

	// initialise the stream
	stream, err := client.StreamServerTime(ctx, &proto.StreamServerTimeRequest{
		IntervalSeconds: 2,
	})

	if err != nil {
		log.Fatal(err)
	}

	// go monitorConnectionState(conn)

	// loop through all the responses we get back from the server
	// - log each response
	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		log.Printf("received time from server: %s", res.CurrentTime.AsTime())
	}

	// once the server closes the stream exit gracefully
	log.Println("server stream closed")
}

// NOT NEEDED, but cool to know:
// func monitorConnectionState(conn *grpc.ClientConn) {
// 	for {
// 		state := conn.GetState()
// 		fmt.Println("Connection state:", state)
//
// 		if state == connectivity.TransientFailure || state == connectivity.Shutdown {
// 			log.Println("Detected server shutdown or connectivity issue")
// 			break
// 		}
//
// 		// Wait until the state changes
// 		conn.WaitForStateChange(context.Background(), state)
// 	}
// }
