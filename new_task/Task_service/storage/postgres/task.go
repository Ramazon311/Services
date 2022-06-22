package postgres

import (
	"database/sql"
	pb "github/Services/newpro/Task_service/genproto/task_service"
	"time"

	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Create(Task *pb.Task) (*pb.Task, error) {
	query := `
        INSERT INTO 
			task1 (id,name, title, assignee_id, summary, deadline, status, created_at)
        VALUES($1,$2,$3,$4,$5,$6,$7,$8)
        RETURNING id
    `
	Task.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := r.db.DB.QueryRow(query,
		Task.Id,
		Task.Name,
		Task.Title,
		Task.AssigneeId,
		Task.Summary,
		Task.Deadline,
		Task.Status,
		Task.CreatedAt,
	).Scan(&Task.Id)
	if err != nil {
		return nil, err
	}
	return r.Get(Task.Id)
}

func (r *TaskRepo) Get(id string) (*pb.Task, error) {
	query := `
        SELECT
            name, assignee_id, title, summary, deadline, status, created_at
        FROM task1
        WHERE id = $1
		AND deleted_at IS NULL
    `
	var task pb.Task
	err := r.db.DB.QueryRow(query,
		id,
	).Scan(
		&task.Name,
		&task.AssigneeId,
		&task.Title,
		&task.Summary,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepo) List(req *pb.ListReq) (*pb.ListRes, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListRes
	query := `
		SELECT 
			name, assignee_id, title, summary, deadline, status, created_at
        FROM task1
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Name,
			&task.AssigneeId,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	query = `
		SELECT count(*) FROM task1
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *TaskRepo) Update(Task *pb.Task) (*pb.Task, error) {
	query := `
		UPDATE task1 SET
            assignee_id = $1, title = $2, summary = $3, deadline = $4, status = $5, updated_at = $6, name = $7
        WHERE id = $8
		AND deleted_at IS NULL
    `
	Task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	result, err := r.db.DB.Exec(query,
		Task.AssigneeId,
		Task.Title,
		Task.Summary,
		Task.Deadline,
		Task.Status,
		Task.UpdatedAt,
		Task.Name,

		Task.Id,
	)
	if err != nil {
		return nil, err
	}

	i, _ := result.RowsAffected()
	if i == 0 {
		return nil, sql.ErrNoRows
	}

	task, err := r.Get(Task.Id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepo) Delete(id *pb.IdReq) (*pb.EmptyResp, error) {
	query := `
		UPDATE task1
		SET 
			deleted_at = $1
		WHERE id = $2
	`
	newTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.DB.Exec(query, newTime, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.EmptyResp{}, nil
}

func (r *TaskRepo) ListOverdue(req *pb.ListOverReq) (*pb.ListOverResp, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListOverResp
	query := `
		SELECT 
			name, assignee_id, title, summary, deadline, status,	created_at
		FROM task1
		WHERE deadline < $1
		AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3 `
	rows, err := r.db.DB.Query(query, req.Time, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Name,
			&task.AssigneeId,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	query = `
		SELECT count(*) FROM task1 
		WHERE deadline < $1
		AND deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query, req.Time).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *TaskRepo) GetList(req *pb.Aid) (*pb.ListResp, error) {
	var resp pb.ListResp
	query := `
		SELECT 
			name, assignee_id, title, summary, deadline, status, created_at
        FROM task1
		WHERE deleted_at IS NULL and assignee_id = $1
	`
	rows, err := r.db.DB.Query(query, req.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(
			&task.Name,
			&task.AssigneeId,
			&task.Title,
			&task.Summary,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Tasks = append(resp.Tasks, &task)
	}
	
	return &resp, nil
}
