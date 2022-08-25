package storage

import (
	"github.com/Services/imanuz_service/data_service/storage/postgres"
	"github.com/Services/imanuz_service/data_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

//IStorage ...
type IStorage interface {
	Data() repo.DataStorageI
}

type storagePg struct {
	db       *sqlx.DB
	DataRepo repo.DataStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		DataRepo: postgres.NewDataRepo(db),
	}
}

func (s storagePg) Data() repo.DataStorageI {
	return s.DataRepo
}
