package postgres

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"

	pb "github.com/Gorilla-services/user-service1/genproto"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	
	var (
		rUser = pb.User{}
	)
    ID,err :=uuid.NewV4()
	if err !=nil{
		return nil,err
	}
	err = r.db.QueryRow("INSERT INTO users1 (id, first_name, last_name) VALUES($1, $2, $3) RETURNING id, first_name, last_name", ID, user.FirstName, user.LastName).Scan(
		&rUser.Id,
		&rUser.FirstName,
		&rUser.LastName,
	)
	
	if err != nil {
		return &pb.User{}, err
	}
    fmt.Print(rUser)
	return &rUser, nil
}

func (r *userRepo) UpdateUser(user *pb.User) (string, error) {
	var (
		rUser = pb.User{}
	)

	err := r.db.QueryRow("INSERT INTO users1 (id, first_name, last_name) VALUES($1, $2, $3) RETURNING id, first_name, last_name", user.Id, user.FirstName, user.LastName).Scan(
		&rUser.Id,
		&rUser.FirstName,
		&rUser.LastName,
	)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (r *userRepo) GetByUserId(ID string) (*pb.User, error) {
	var (
		rUser = pb.User{}
	)

	err := r.db.QueryRow("SELECT id, first_name, last_name from users1 WHERE id = $1", ID).Scan(
		&rUser.Id,
		&rUser.FirstName,
		&rUser.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &rUser, nil
}
func (r *userRepo) GetUsersList(page, limit int64)([]*pb.User,int64,error){

	offset := (page-1)*limit
	var (
	//	rUser = pb.GetUsersResponce{}
		users = []*pb.User{}
		count int64
	)
    query := "SELECT * from users1  OFFSET $1 LIMIT $2"
	fmt.Print(page,limit)
	rows,err := r.db.Query(query,offset,limit)
	if err != nil {
		return nil, 0,err
	}
	for rows.Next(){
		var user pb.User
		err :=rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)
		if err!=nil{
			return nil,0,err
		}
		users=append(users, &user)

	}
	countQuery := "SELECT count(*) FROM users1"
	err = r.db.QueryRow(countQuery).Scan(&count)
	if err != nil{
		return nil,0,err
	}
   // fmt.Print(users)
	return users,count, nil
}
