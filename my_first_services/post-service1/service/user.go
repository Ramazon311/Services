package service

import (
	"context"
	"fmt"

	pb "github.com/Gorilla-services/post-service1/genproto"
	l "github.com/Gorilla-services/post-service1/pkg/logger"
	"github.com/Gorilla-services/post-service1/storage"
	//"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	s.logger.Error("failed while generating uuid for new post", l.Error(err))
	// 	return nil, status.Error(codes.Internal, "failed while generating uuid")
	// }
	// req.Id = id.String()
	user, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("failed while inserting post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while inserting post")
	}

	return user, nil
}

func (s *PostService) GetById(ctx context.Context, req *pb.GetByUserIdRequest) (*pb.Post, error) {
	user, err := s.storage.Post().GetById(req.UserId)
	if err != nil {
		s.logger.Error("failed get post", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get post")
	}
    fmt.Println(user)
	return user, err
}

func (s *PostService) GetAllUserPosts(ctx context.Context, req *pb.GetUserPostsrequest) (*pb.GetUserPosts, error) {
	posts, err := s.storage.Post().GetAllUserPosts(req.UserId)
	if err != nil {
		s.logger.Error("failed get all user posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed get all user posts")
	}

	return &pb.GetUserPosts{
		Posts: posts,
	}, err
}
