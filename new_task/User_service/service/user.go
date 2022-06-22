package service

import (
	"context"
	"fmt"

	pbt "github/Services/newpro/User_service/genproto/task_service"
	pb "github/Services/newpro/User_service/genproto/user_service"
	l "github/Services/newpro/User_service/pkg/logger"
	grpcClient "github/Services/newpro/User_service/service/grpc_client"
	"github/Services/newpro/User_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.Login1) (*pb.User, error) {

	User, err := s.storage.User().Login(req)
	if err != nil {
		s.logger.Error("Error getting User Profile \n \n", l.Error(err))
		return nil, err
	}

	return User, nil
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {

	req.Id = uuid.New().String()
	User, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error create User ", l.Error(err))
		return nil, err
	}
	fmt.Println(" error", err)
	return User, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.IdReq) (*pb.User, error) {

	User, err := s.storage.User().Get(req.Id)
	if err != nil {
		s.logger.Error("Error get User", l.Error(err))
		return nil, err
	}

	Task, err := s.client.TaskService().GetList(ctx, &pbt.Aid{
		Id: req.Id,
	})
	fmt.Println(User.Id)
	if err != nil {
		s.logger.Error("Error get User's tasks", l.Error(err))
		return nil, err
	}

	for _, val := range Task.Tasks {

		User.Task = append(User.Task, &pb.Task{
			AssigneeId: val.AssigneeId,
			Name:       val.Name,
			Deadline:   val.Deadline,
			Summary:    val.Summary,
			Status:     val.Status,
			Title:      val.Title,
			CreatedAt:  val.CreatedAt,
		})
	}

	return User, nil
}

func (s *UserService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	resp, err := s.storage.User().List(req)
	if err != nil {
		s.logger.Error("Error list Users", l.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
	User, err := s.storage.User().Update(req)
	if err != nil {
		s.logger.Error("Error update User", l.Error(err))
		return nil, err
	}
	return User, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.IdReq) (*pb.EmptyResp, error) {
	_, err := s.storage.User().Delete(req)
	if err != nil {
		s.logger.Error("Error delete User", l.Error(err))
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}

func (s *UserService) CheckField(ctx context.Context, req *pb.Check) (*pb.Status, error) {

	bol, err := s.storage.User().CheckField(req)

	if err != nil {
		s.logger.Error("Error delete User", l.Error(err))
		return &pb.Status{Status: false}, err
	}

	return bol, nil
}
