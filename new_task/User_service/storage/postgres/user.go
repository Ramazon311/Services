package postgres

import (
	"database/sql"
	"fmt"
	pb "github/Services/newpro/User_service/genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(User *pb.User) (*pb.User, error) {
	query := `
        INSERT INTO 
			user1 (id,first_name, 
				   last_name, username, 
				   profile_photo, bio, 
				   email, gender, 
				   address, phone_number, 
				   created_at, password, 
				   access_token, refresh_token)
        VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
        RETURNING id
    `
	User.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := r.db.DB.QueryRow(query,
		User.Id,
		User.FirstName,
		User.LastName,
		User.Username,
		User.ProfilePhoto,
		User.Bio,
		User.Email,
		User.Gender,
		User.Address,
		User.PhoneNumber,
		User.CreatedAt,
		User.Password,
		User.AccessToken,
		User.RefreshToken,
	).Scan(&User.Id)
	if err != nil {
		return nil, err
	}
	return r.Get(User.Id)
}

func (r *UserRepo) Login(L *pb.Login1) (*pb.User, error) {
	query := `
        SELECT
			first_name, last_name, 
			username, profile_photo, 
			bio, email, gender, address, 
			phone_number, created_at,
			password, 
			access_token, refresh_token
        FROM user1
        WHERE email = $1
		AND deleted_at IS NULL
    `
	var User pb.User
	err := r.db.DB.QueryRow(query,
		L.Email,
	).Scan(
		&User.FirstName,
		&User.LastName,
		&User.Username,
		&User.ProfilePhoto,
		&User.Bio,
		&User.Email,
		&User.Gender,
		&User.Address,
		&User.PhoneNumber,
		&User.CreatedAt,
		&User.Password,
		&User.AccessToken,
		&User.RefreshToken,
	)
	if err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *UserRepo) Get(id string) (*pb.User, error) {
	query := `
        SELECT
			first_name, last_name, 
			username, profile_photo, 
			bio, email, gender, address, 
			phone_number, created_at,
			password, 
			access_token, refresh_token
        FROM user1
        WHERE id = $1
		AND deleted_at IS NULL
    `
	var User pb.User
	err := r.db.DB.QueryRow(query,
		id,
	).Scan(
		&User.FirstName,
		&User.LastName,
		&User.Username,
		&User.ProfilePhoto,
		&User.Bio,
		&User.Email,
		&User.Gender,
		&User.Address,
		&User.PhoneNumber,
		&User.CreatedAt,
		&User.Password,
		&User.AccessToken,
		&User.RefreshToken,
	)
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *UserRepo) List(req *pb.ListReq) (*pb.ListResp, error) {
	offset := (req.Page - 1) * req.Limit
	var resp pb.ListResp
	query := `
		SELECT 
			first_name, last_name, 
			username, profile_photo, 
			bio, email, gender, address, 
			phone_number, created_at
        FROM user1
		WHERE deleted_at IS NULL
		LIMIT $1
		OFFSET $2
	`
	rows, err := r.db.DB.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var User pb.User
		err = rows.Scan(
			&User.FirstName,
			&User.LastName,
			&User.Username,
			&User.ProfilePhoto,
			&User.Bio,
			&User.Email,
			&User.Gender,
			&User.Address,
			&User.PhoneNumber,
			&User.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &User)
	}
	query = `
		SELECT count(*) FROM user1
		WHERE deleted_at IS NULL
	`
	err = r.db.DB.QueryRow(query).Scan(&resp.Count)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *UserRepo) Update(User *pb.User) (*pb.User, error) {
	query := `
		UPDATE user1 SET
			first_name = $1, last_name = $2, 
			username = $3, profile_photo = $4, 
			bio = $5, email = $6, gender = $7, address = $8, 
			phone_number = $9
		WHERE id = $10
		AND deleted_at IS NULL
        RETURNING id
    `
	User.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	result, err := r.db.DB.Exec(query,
		User.FirstName,
		User.LastName,
		User.Username,
		User.ProfilePhoto,
		User.Bio,
		User.Email,
		User.Gender,
		User.Address,
		User.PhoneNumber,
		User.Id,
	)
	if err != nil {
		return nil, err
	}

	i, _ := result.RowsAffected()
	if i == 0 {
		return nil, sql.ErrNoRows
	}

	user, err := r.Get(User.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Delete(id *pb.IdReq) (*pb.EmptyResp, error) {
	query := `
		UPDATE user1
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

func (r *UserRepo) CheckField(req *pb.Check) (*pb.Status, error) {
	var existsClient int

	if req.Field == "username" {

		row := r.db.QueryRow(`
			SELECT count(1) FROM user1 WHERE  username = $1 AND deleted_at IS NULL`, req.Value,
		)
		if err := row.Scan(&existsClient); err != nil {
			return &pb.Status{Status: false}, err
		}

	} else if req.Field == "email" {

		row := r.db.QueryRow(`
			SELECT count(1) FROM user1 WHERE  email = $1 AND deleted_at IS NULL`, req.Value,
		)
		if err := row.Scan(&existsClient); err != nil {
			return &pb.Status{Status: false}, err
		}

	} else {
		return &pb.Status{Status: false}, nil
	}
	fmt.Println("\n\n\n", existsClient)
	if existsClient == 0 {
		return &pb.Status{Status: false}, nil
	}

	return &pb.Status{Status: true}, nil
}
