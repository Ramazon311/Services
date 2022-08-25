package grpc_client

import (
	"fmt"
	"github/Services/newpro/User_service/config"

	pb "github/Services/newpro/User_service/genproto/task_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	TaskService() pb.TaskServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	taskService pb.TaskServiceClient
}

func New(cfg config.Config) (*GrpcClient, error) {

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.TaskServiceHost, cfg.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
			cfg.TaskServiceHost, cfg.TaskServicePort)
	}

	grpcClient := &GrpcClient{
		cfg:         cfg,
		taskService: pb.NewTaskServiceClient(connTask),
	}

	return grpcClient, nil
}

// TaskService ...
func (c *GrpcClient) TaskService() pb.TaskServiceClient {
	return c.taskService
}
