package repo

import (
	pb "github/Services/post_task/data_service/genproto/data_service"
)

//PostStorageI ...

type DataStorageI interface {
	Create(*pb.Post) (*pb.EmptyResp, error)
}


