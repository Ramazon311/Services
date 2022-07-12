package services

import (
	"fmt"

	"github.com/Services/imanuz_service/api_service/config"
	pbu "github.com/Services/imanuz_service/api_service/genproto/data_service"
	pb "github.com/Services/imanuz_service/api_service/genproto/post_service"

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
