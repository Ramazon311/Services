package repo

import (
	pb "github.com/Gorilla-services/user-service1/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	UpdateUser(*pb.User) (string, error)
	GetByUserId(id string) (*pb.User, error)
	GetUsersList(limit,page int64)  (  []*pb.User,int64,error)
}
