package repo

import (
	pb "github.com/Services/imanuz_service/data_service/genproto/data_service"
)

//PostStorageI ...

type DataStorageI interface {
	Create(*pb.Post) (*pb.EmptyResp, error)
}
