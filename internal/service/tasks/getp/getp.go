package getp

import (
	"context"
	"fmt"

	"github.com/pyramidum-space/protos/gen/go/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	client tasks.TasksServiceClient
}

func NewService(serverAddress string) (*Service, error) {
	const op = "services.tasks.getp.NewService"

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client := tasks.NewTasksServiceClient(conn)
	return &Service{
		client: client,
	}, nil
}

func (s *Service) Get(task_id []byte) (*tasks.Task, error) {
	const op = "services.tasks.post.Create"

	response, err := s.client.Get(context.TODO(), &tasks.GetRequest{
		TaskId: task_id,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response.Task, nil
}
