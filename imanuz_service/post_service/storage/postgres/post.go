package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Services/imanuz_service/post_service/genproto/post_service"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) Get(id string) (*pb.Post, error) {
	query := `
        SELECT
            id, user_id, title, body, updated_at, created_at
        FROM post7
        WHERE id = $1
		AND deleted_at IS NULL
    `
	var (
		post       pb.Post
		nullupdate sql.NullString
	)
	err := r.db.DB.QueryRow(query,
		id,
	).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
		&nullupdate,
		&post.CreatedAt,
	)
	if nullupdate.Valid {
		post.UpdatedAt = nullupdate.String
	}
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepo) List(req *pb.ListReq) (*pb.ListRes, error) {
	offset := (req.Page - 1) * req.Limit
	fmt.Println("\n\n>>", req)
	var (
		resp       pb.ListRes
		nullupdate sql.NullString
	)
	query := `
		SELECT 
			id, user_id, title, body, updated_at, created_at
        FROM post7
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post pb.Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Body,
			&nullupdate,
			&post.CreatedAt,
		)
		if nullupdate.Valid {
			post.UpdatedAt = nullupdate.String
		}
		if err != nil {
			return nil, err
		}

		resp.Posts = append(resp.Posts, &post)
	}
	query = `
		SELECT count(*) FROM post7
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *PostRepo) Update(Task *pb.UpdatePost) (*pb.Post, error) {
	query := `
		UPDATE post7 SET
            title = $1, body = $2, updated_at = $3
        WHERE id = $4
		AND deleted_at IS NULL
    `
	updated_at := time.Now().Format("2006-01-02 15:04:05")

	result, err := r.db.DB.Exec(query,
		Task.Title,
		Task.Body,
		updated_at,
		Task.Id,
	)
	if err != nil {
		return nil, err
	}

	i, _ := result.RowsAffected()
	if i == 0 {
		return nil, sql.ErrNoRows
	}
	post, err := r.Get(Task.Id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepo) Delete(id string) (*pb.EmptyResp, error) {
	query := `
		UPDATE post7
		SET 
			deleted_at = $1
		WHERE id = $2
	`
	newTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.DB.Exec(query, newTime, id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}
