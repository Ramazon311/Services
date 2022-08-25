package service

import (
	"context"
	"fmt"

	em "github/Services/newpro/Task_service/genproto/email_service"
	pb "github/Services/newpro/Task_service/genproto/task_service"
	pbt "github/Services/newpro/Task_service/genproto/user_service"

	l "github/Services/newpro/Task_service/pkg/logger"
	grpcClient "github/Services/newpro/Task_service/service/grpc_client"
	"github/Services/newpro/Task_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.GrpcClientI
}

func NewTaskService(db *sqlx.DB, log l.Logger, client grpcClient.GrpcClientI) *TaskService {
	return &TaskService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	req.Id = uuid.New().String()
	task, err := s.storage.Task().Create(req)
	if err != nil {
		s.logger.Error("Error create task \n \n", l.Error(err))
		return nil, err
	}
	fmt.Println("\n\n\n error", err)
	return task, nil
}

func (s *TaskService) Get(ctx context.Context, req *pb.IdReq) (*pb.Task, error) {
	task, err := s.storage.Task().Get(req.Id)
	if err != nil {
		s.logger.Error("Error get task", l.Error(err))
		return nil, err
	}
	// user, err := s.client.UserService().Get(ctx, &pbu.IdReq{
	// 	Id: task.AssigneeId,
	// })
	// task.User.Id = user.Id
	// task.User.FirstName = user.FirstName
	// task.User.LastName = user.LastName
	// task.User.Username = user.Username
	// task.User.Email = user.Email
	// task.User.Address = user.Address
	// task.User.Bio = user.Bio
	// task.User.CreatedAt = user.Bio
	// task.User.DeletedAt = user.DeletedAt
	// task.User.UpdatedAt = user.UpdatedAt
	// task.User.Gender = user.Gender
	// task.User.PhoneNumber = user.PhoneNumber
	// task.User.ProfilePhoto = user.ProfilePhoto

	// if err != nil {
	// 	s.logger.Error("Error get user", l.Error(err))
	// 	return nil, err
	// }

	return task, nil
}

func (s *TaskService) List(ctx context.Context, req *pb.ListReq) (*pb.ListRes, error) {
	resp, err := s.storage.Task().List(req)
	if err != nil {
		s.logger.Error("Error list tasks", l.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *TaskService) GetList(ctx context.Context, req *pb.Aid) (*pb.ListResp, error) {
	resp, err := s.storage.Task().GetList(req)
	fmt.Println(resp)
	if err != nil {
		s.logger.Error("Error list tasks", l.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *TaskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Update(req)
	if err != nil {
		s.logger.Error("Error update task", l.Error(err))
		return nil, err
	}
	return task, nil
}

func (s *TaskService) Delete(ctx context.Context, req *pb.IdReq) (*pb.EmptyResp, error) {
	_, err := s.storage.Task().Delete(req)
	if err != nil {
		s.logger.Error("Error delete task", l.Error(err))
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}

func (s *TaskService) ListOverdue(ctx context.Context, req *pb.ListOverReq) (*pb.ListOverResp, error) {
	tasks, err := s.storage.Task().ListOverdue(req)
	if err != nil {
		s.logger.Error("Error list overdue task", l.Error(err))
		return nil, err
	}

	var emails = []string{}
	var phones = []string{}

	for _, val := range tasks.Tasks {
		User, err := s.client.UserService().Get(ctx, &pbt.IdReq{Id: val.AssigneeId})
		if err != nil {
			s.logger.Error("Error get User in get ListOverdue", l.Error(err))
			return nil, err
		}
		emails = append(emails, User.Email)
		phones = append(phones, User.PhoneNumber)
	}

	info := &em.EmailText{
		Id:        "",
		Subject:   "Task",
		Body:      "Your limit time finishied",
		Phone:     phones,
		Recipints: emails,
	}

	_, err = s.client.EmailService().Send(ctx, info)

	if err != nil {
		s.logger.Error("Sending email for User", l.Error(err))
		return nil, err
	}

	return tasks, nil
}
