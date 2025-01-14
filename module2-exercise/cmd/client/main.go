package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/alessioalex/grpc-course/module2-exercise/proto"
)

func main() {
	ctx := context.Background()

	// conn, err := grpc.Dial(
	// 	"localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithBlock(),
	// )
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	todos := []string{"Walk the dog", "Groceries", "Pick up kid" /*, ""*/}

	for _, todo := range todos {
		res, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: todo})
		if err != nil {
			s, ok := status.FromError(err)
			if ok {
				log.Fatalf("status code: %s, error: %s", s.Code().String(), s.Message())
			}
			log.Fatal(err)
		}

		log.Printf("response received: %s", res.GetId())
	}

	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: "2"})
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			log.Fatalf("status code: %s, error: %s", s.Code().String(), s.Message())
		}
		log.Fatal(err)
	}

	res, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", res.GetTasks())

}
