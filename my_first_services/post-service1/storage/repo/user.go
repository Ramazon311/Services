package repo

import (
	pb "github.com/Gorilla-services/post-service1/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	Create(*pb.Post) (*pb.Post, error)
	GetById(id string) (*pb.Post, error)
	GetAllUserPosts(userID string) ([]*pb.Post, error)
}
