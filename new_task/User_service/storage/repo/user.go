package repo

import (
	pb "github/Services/newpro/User_service/genproto/user_service"
)

//UserStorageI ...
type UserStorageI interface {
	Create(*pb.User) (*pb.User, error)
	Get(id string) (*pb.User, error)
	List(*pb.ListReq) (*pb.ListResp, error)
	Update(*pb.User) (*pb.User, error)
	Delete(*pb.IdReq) (*pb.EmptyResp, error)
	CheckField(*pb.Check) (*pb.Status, error)
	Login(*pb.Login1) (*pb.User, error)
}
