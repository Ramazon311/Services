package repo

import (
	pb "github.com/Services/imanuz_service/post_service/genproto/post_service"
)

//PostStorageI ...

type PostStorageI interface {
	Get(string) (*pb.Post, error)
	List(*pb.ListReq) (*pb.ListRes, error)
	Update(*pb.UpdatePost) (*pb.Post, error)
	Delete(string) (*pb.EmptyResp, error)
}
