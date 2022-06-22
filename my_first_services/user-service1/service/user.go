package service

import (
	"context"
	"fmt"

	pb "github.com/Gorilla-services/user-service1/genproto"
	l "github.com/Gorilla-services/user-service1/pkg/logger"
	cl "github.com/Gorilla-services/user-service1/service/grpc_client"
	"github.com/Gorilla-services/user-service1/storage"
//	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	s.logger.Error("failed while generating uuid for new user", l.Error(err))
	// 	return nil, status.Error(codes.Internal, "failed while generating uuid")
	// }
	// req.Id = id.String()
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		fmt.Println(user)
		s.logger.Error("failed while inserting user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while inserting user")
	}

	for _, post := range user.Posts {
		post.UserId = req.Id
	    createdPosts ,err := s.client.PostService().Create(ctx,post)
		//createdPosts, err := s.client.PostService().Create(ctx, post)
		if err != nil {
			fmt.Println("dd",createdPosts,"dd")
			s.logger.Error("failed while inserting user post", l.Error(err))
			return nil, status.Error(codes.Internal, "failed while inserting user post")
		}
		fmt.Println(createdPosts)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.UpdateUserResponse, error) {
	id, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error("failed while inserting user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while inserting user")
	}

	return &pb.UpdateUserResponse{
		Id: id,
	}, nil
}

func (s *UserService) GetByUserId(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	user, err := s.storage.User().GetByUserId(req.UserId)
	if err != nil {
		s.logger.Error("failed get user", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get user")
	}

	userPosts, err := s.client.PostService().GetAllUserPosts(ctx, &pb.GetUserPostsrequest{
		UserId: req.UserId,
	})
	if err != nil {
		//fmt.Print(userPosts)
		s.logger.Error("failed get user posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get user posts")
	}

	user.Posts = userPosts.Posts

	return user, err
}
func (s *UserService) GetUserList(ctx context.Context,req *pb.GetUsersRequest)(*pb.GetUsersResponce,error){

	users,count,err := s.storage.User().GetUsersList(req.Limit,req.Page)
	if err !=nil{
		s.logger.Error("failed get users", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get users")
	}
	
	for _,user := range users{
		post,err :=s.client.PostService().GetAllUserPosts(ctx,&pb.GetUserPostsrequest{
			UserId : user.Id,
		})
		if err!= nil{
			s.logger.Error("failed get user posts", l.Error(err))
	      	return nil, status.Error(codes.Internal, "failed get user posts")
		}
		user.Posts=post.Posts
		//users.User1=append(users.User1, user)
	}
	return &pb.GetUsersResponce{
		User1: users,
		Count: count,
	},nil

}
