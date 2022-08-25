package repo

import (
	pb "github/Services/newpro/Task_service/genproto/task_service"
	em "github/Services/newpro/Task_service/genproto/email_service"
)

//TaskStorageI ...

type TaskStorageI interface {
	Create(*pb.Task) (*pb.Task, error)
	Get(string) (*pb.Task, error)
	List(*pb.ListReq) (*pb.ListRes, error)
	GetList(*pb.Aid) (*pb.ListResp, error)
	Update(*pb.Task) (*pb.Task, error)
	Delete(*pb.IdReq) (*pb.EmptyResp, error)
	ListOverdue(*pb.ListOverReq) (*pb.ListOverResp, error)
}


//EmailStorageI

type EmailStorageI interface {
	Email(*em.EmailText)(*em.Empty, error)
}