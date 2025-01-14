package todo

import (
	"context"
	"unicode/utf8"

	"github.com/alessioalex/grpc-course/module2-exercise/internal/tasks"
	"github.com/alessioalex/grpc-course/module2-exercise/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	proto.UnimplementedTodoServiceServer
	list tasks.List
}

func NewService() *Service {
	return &Service{
		list: *tasks.NewList(),
	}
}

// - AddTask
//   - Will add a task to the todo map
//   - AddTaskRequest message contains a task string
//   - AddTaskResponse message contains a generated task ID
//   - Server should return InvalidArgument if task is empty
func (s *Service) AddTask(
	ctx context.Context,
	request *proto.AddTaskRequest,
) (*proto.AddTaskResponse, error) {
	task := request.GetTask()

	if utf8.RuneCountInString(task) < 4 {
		return nil, status.Error(
			codes.InvalidArgument,
			"task name cannot be less than 4 characters",
		)
	}

	id := s.list.Add(task)

	return &proto.AddTaskResponse{Id: id}, nil
}

// - CompleteTask
//   - Will remove a task from the todo map
//   - CompleteTaskRequest message contains a task ID
//   - CompleteTaskResponse is an empty response
//   - Server should return InvalidArgument if task ID is empty
//   - Server should return NotFound if task ID is not found
func (s *Service) CompleteTask(
	ctx context.Context,
	request *proto.CompleteTaskRequest,
) (*proto.CompleteTaskResponse, error) {
	id := request.GetId()

	if id == "" {
		return nil, status.Error(
			codes.InvalidArgument,
			"task id cannot be empty",
		)
	}

	ok := s.list.Del(id)
	if !ok {
		return nil, status.Error(
			codes.NotFound,
			"task ID not found",
		)
	}

	return &proto.CompleteTaskResponse{}, nil
}

// - ListTasks
//   - Will return a list of outstanding tasks
//   - ListTasksRequest message is an empty request
//   - ListTasksResponse message contains a list of outstanding tasks
func (s *Service) ListTasks(
	ctx context.Context,
	request *proto.ListTasksRequest,
) (*proto.ListTasksResponse, error) {
	// var tasks []*proto.Task
	taskList := s.list.All()
	tasks := make([]*proto.Task, len(taskList))

	for id, task := range taskList {
		tasks = append(tasks, &proto.Task{Id: id, Task: task})
	}

	return &proto.ListTasksResponse{Tasks: tasks}, nil
}
