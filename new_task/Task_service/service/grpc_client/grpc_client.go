package grpc_client

import (
	"fmt"
	"github/Services/newpro/Task_service/config"
	pb "github/Services/newpro/Task_service/genproto/user_service"
	em "github/Services/newpro/Task_service/genproto/email_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	UserService() pb.UserServiceClient
	EmailService() em.EmailServiceClient 
}

type GrpcClient struct {
	cfg         config.Config
	userService pb.UserServiceClient
	emailServie em.EmailServiceClient
}

func New(cfg config.Config) (*GrpcClient, error) {
	
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port: %d",
			cfg.UserServiceHost, cfg.UserServicePort)
	}

	grpcClient := &GrpcClient{
		cfg:         cfg,
		userService: pb.NewUserServiceClient(connUser),
	}

	return grpcClient, nil
}

// UserService ...
func (c *GrpcClient) UserService() pb.UserServiceClient {
	return c.userService
}

// EmailService ...
func (c *GrpcClient) EmailService() em.EmailServiceClient {
	return c.emailServie
}