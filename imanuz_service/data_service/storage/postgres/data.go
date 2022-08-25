package postgres

import (
	"time"

	pb "github.com/Services/imanuz_service/data_service/genproto/data_service"

	"github.com/jmoiron/sqlx"
)

type DataRepo struct {
	db *sqlx.DB
}

func NewDataRepo(db *sqlx.DB) *DataRepo {
	return &DataRepo{db: db}
}

func (r *DataRepo) Create(post *pb.Post) (*pb.EmptyResp, error) {

	query := `INSERT INTO post7 (id, user_id, title, body, created_at)
			VALUES($1,$2,$3,$4,$5)`

	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := r.db.DB.QueryRow(query,
		post.Id,
		post.UserId,
		post.Title,
		post.Body,
		post.CreatedAt,
	)

	if err != nil {
		return nil, nil
	}

	return &pb.EmptyResp{}, nil
}
