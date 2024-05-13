package register

import (
	"context"
	"fmt"

	"github.com/pyramidum-space/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	client auth.AuthClient
}

func NewService(serverAddress string) (*Service, error) {
	const op = "services.auth.register.Register"

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	client := auth.NewAuthClient(conn)
	return &Service{
		client: client,
	}, nil
}

func (s *Service) Register(email string, password string) (int64, error) {
	const op = "services.auth.register.Register"

	response, err := s.client.Register(context.TODO(), &auth.RegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return response.UserId, nil
}
