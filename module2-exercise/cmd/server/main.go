package main

import (
	"log"
	"net"

	"github.com/alessioalex/grpc-course/module2-exercise/internal/todo"
	"github.com/alessioalex/grpc-course/module2-exercise/proto"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	todoService := todo.NewService()

	proto.RegisterTodoServiceServer(grpcServer, todoService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error spinning up grpc server: %v", err)
	}

	log.Printf("Starting gRPC server on address: %s", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving grpc server: %v", err)
	}
}
