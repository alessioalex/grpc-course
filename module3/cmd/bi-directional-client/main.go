package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/alessioalex/grpc-course/module3/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// initialise our grpc connection
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
	stream, err := client.Echo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// create a separate goroutine to listen to the server responses
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		// loop for each message from server
		for {
			res, err := stream.Recv()
			if err != nil {
				// check if stream is closed
				if err == io.EOF {
					// close the client stream
					return nil
				}
				return err
			}

			// log the message
			log.Printf("message received from server: %s", res.GetMessage())
		}
	})

	// send some log messages
	for i := range 5 {
		req := &proto.EchoRequest{
			Message: fmt.Sprintf("Foobar: %d", i),
		}

		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * 2)
	}

	// close the client stream
	if err := stream.CloseSend(); err != nil {
		log.Fatal(err)
	}

	// wait for the server goroutine to finish
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("bi-directional stream closed")
}
