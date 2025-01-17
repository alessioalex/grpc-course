package streaming

import (
	"io"
	"log"
	"time"

	"github.com/alessioalex/grpc-course/module3/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	proto.UnimplementedStreamingServiceServer
}

func (service *Service) StreamServerTime(
	request *proto.StreamServerTimeRequest,
	stream grpc.ServerStreamingServer[proto.StreamServerTimeResponse],
) error {
	// initialize a ticker for our interval
	if request.GetIntervalSeconds() == 0 {
		return status.Error(codes.InvalidArgument, "interval must be set")
	}

	interval := time.Duration(request.GetIntervalSeconds() * int32(time.Second))
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	ticks := 10

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			if ticks == 0 {
				return nil
			}

			ticks--
			currentTime := time.Now()
			resp := &proto.StreamServerTimeResponse{
				CurrentTime: timestamppb.New(currentTime),
			}

			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func (s *Service) LogStream(
	stream grpc.ClientStreamingServer[proto.LogStreamRequest, proto.LogStreamResponse],
) error {
	// initialise a count
	count := 0

	for {
		// receive our message
		logEntry, err := stream.Recv()
		if err != nil {
			// check if the stream is closed
			if err == io.EOF {
				return stream.SendAndClose(&proto.LogStreamResponse{
					EntriesLogged: int32(count),
				})
			}
			return err
		}

		// log message
		log.Printf(
			"Received log: [%s]: %s %s",
			logEntry.GetTimestamp().AsTime(),
			logEntry.GetLevel().String(),
			logEntry.GetMessage(),
		)

		// increment count
		count++
	}
}

func (s *Service) Echo(
	stream grpc.BidiStreamingServer[proto.EchoRequest, proto.EchoResponse],
) error {
	// loop through the messages received from the client
	for {
		req, err := stream.Recv()
		if err != nil {
			// check if the stream is closed
			if err == io.EOF {
				// close the server side stream
				return nil
			}
			return err
		}

		log.Printf("message received: %s", req.GetMessage())

		// - build our response and send back from server
		res := &proto.EchoResponse{
			Message: req.GetMessage(),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
