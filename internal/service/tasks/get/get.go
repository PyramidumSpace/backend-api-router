package get

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
	const op = "services.tasks.get.NewService"

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client := tasks.NewTasksServiceClient(conn)
	return &Service{
		client: client,
	}, nil
}

func (s *Service) Get(user_id int32,
	search string,
	deadLineFrom time.Time,
	deadLineTo time.Time,
	possibleDeadLineFrom time.Time,
	possibleDeadLineTo time.Time,
	progressStatus tasks.ProgressStatus,
	isUrgent bool,
	isImportant bool,
	weightFrom int32,
	weightTo int32) ([]*tasks.Task, error) {
	const op = "services.tasks.post.Create"

	response, err := s.client.Tasks(context.TODO(), &tasks.TasksRequest{
		OwnerId:              user_id,
		Search:               &search,
		DeadlineFrom:         timestamppb.New(deadLineFrom),
		DeadlineTo:           timestamppb.New(deadLineTo),
		PossibleDeadlineFrom: timestamppb.New(possibleDeadLineFrom),
		PossibleDeadlineTo:   timestamppb.New(possibleDeadLineTo),
		IsUrgent:             &isUrgent,
		IsImportant:          &isImportant,
		WeightFrom:           &weightFrom,
		WeightTo:             &weightTo,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response.Tasks, nil
}
