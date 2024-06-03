package put

import (
	"context"
	"fmt"
	"time"

	"github.com/pyramidum-space/protos/gen/go/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	client tasks.TasksServiceClient
}

func NewService(serverAddress string) (*Service, error) {
	const op = "services.tasks.put.NewService"

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client := tasks.NewTasksServiceClient(conn)
	return &Service{
		client: client,
	}, nil
}

func (s *Service) Update(id []byte,
	header string,
	text string,
	external_images []string,
	deadline time.Time,
	progress_status tasks.ProgressStatus,
	is_urgent bool,
	is_important bool,
	owner_id int32,
	parent_id []byte,
	possible_deadline time.Time,
	weight int32) ([]byte, error) {
	const op = "services.tasks.post.Create"

	task := tasks.Task{
		Id:               id,
		Header:           header,
		Text:             text,
		ExternalImages:   external_images,
		Deadline:         timestamppb.New(deadline),
		ProgressStatus:   progress_status,
		IsUrgent:         is_urgent,
		IsImportant:      is_important,
		OwnerId:          owner_id,
		ParentId:         parent_id,
		PossibleDeadline: timestamppb.New(possible_deadline),
		Weight:           weight,
	}

	response, err := s.client.Update(context.TODO(), &tasks.UpdateRequest{
		Task: &task,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response.TaskId, nil
}
