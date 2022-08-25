package service

import (
	"context"

	pb "github.com/Services/imanuz_service/post_service/genproto/post_service"

	l "github.com/Services/imanuz_service/post_service/pkg/logger"
	"github.com/Services/imanuz_service/post_service/storage"

	"github.com/jmoiron/sqlx"
)

type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) Get(ctx context.Context, req *pb.IdP) (*pb.Post, error) {
	post, err := s.storage.Post().Get(req.Id)
	if err != nil {
		s.logger.Error("Error get post", l.Error(err))
		return nil, err
	}

	return post, nil
}

func (s *PostService) List(ctx context.Context, req *pb.ListReq) (*pb.ListRes, error) {
	resp, err := s.storage.Post().List(req)
	if err != nil {
		s.logger.Error("Error list posts", l.Error(err))
		return nil, err
	}
	return resp, nil
}

func (s *PostService) Update(ctx context.Context, req *pb.UpdatePost) (*pb.Post, error) {
	post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error("Error update post", l.Error(err))
		return nil, err
	}
	return post, nil
}

func (s *PostService) Delete(ctx context.Context, req *pb.IdP) (*pb.EmptyResp, error) {
	_, err := s.storage.Post().Delete(req.Id)
	if err != nil {
		s.logger.Error("Error delete post", l.Error(err))
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}
