package services

import (
	"fmt"
	"github/Services/newpro/Api-TU/config"
	em "github/Services/newpro/Api-TU/genproto/email_service"
	pbt "github/Services/newpro/Api-TU/genproto/task_service"
	pbu "github/Services/newpro/Api-TU/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	TaskService() pbt.TaskServiceClient
	UserService() pbu.UserServiceClient
	EmailService() em.EmailServiceClient
}

type serviceManager struct {
	taskService pbt.TaskServiceClient
	userService pbu.UserServiceClient
	emailServie em.EmailServiceClient
}

func (s *serviceManager) TaskService() pbt.TaskServiceClient {
	return s.taskService
}
func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}
func (c *serviceManager) EmailService() em.EmailServiceClient {
	return c.emailServie
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaskServiceHost, conf.TaskServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connEmail, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.EmailServiceHost, conf.EmailServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taskService: pbt.NewTaskServiceClient(connTask),
		userService: pbu.NewUserServiceClient(connUser),
		emailServie: em.NewEmailServiceClient(connEmail),
	}

	return serviceManager, nil
}
