package storage

import (
	"github/Services/newpro/Task_service/storage/postgres"
	"github/Services/newpro/Task_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	Task() repo.TaskStorageI
}

type storagePg struct {
	db       *sqlx.DB
	TaskRepo repo.TaskStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		TaskRepo: postgres.NewTaskRepo(db),
	}
}

func (s storagePg) Task() repo.TaskStorageI {
	return s.TaskRepo
}
