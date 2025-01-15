package streaming

import (
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
