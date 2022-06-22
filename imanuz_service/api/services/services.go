package services

import (
	"fmt"
	"github/Services/post_task/api/config"
	pb "github/Services/post_task/api/genproto/post_service"
	pbu "github/Services/post_task/api/genproto/data_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	PostService() pb.PostServiceClient
	DataService() pbu.DataServiceClient
}

type serviceManager struct {
	postService pb.PostServiceClient
	dataService pbu.DataServiceClient
}

func (s *serviceManager) PostService() pb.PostServiceClient {
	return s.postService
}
func (s *serviceManager) DataService() pbu.DataServiceClient {
	return s.dataService
}


func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connData, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.DataServiceHost, conf.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	

	serviceManager := &serviceManager{
		postService: pb.NewPostServiceClient(connPost),
		dataService: pbu.NewDataServiceClient(connData),
	}

	return serviceManager, nil
}
