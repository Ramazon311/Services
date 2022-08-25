package postgres

import (
	"fmt"
	"time"

	"github/Services/newpro/Email_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type sendRepo struct {
	db *sqlx.DB
}

// NewSendRepo ...
func NewSendRepo(db *sqlx.DB) repo.SendStorageI {
	return &sendRepo{db: db}
}

// MakeSent ...

func (s *sendRepo) CreatEmailText(id, subject, body string) (string, error) {
	query := `INSERT INTO email_text(
		id, 
		subject, 
		body, 
		created_at)
		VALUES($1, $2, $3, $4)
		RETURNING id;`

	err := s.db.DB.QueryRow(
		query,
		id,
		subject,
		body,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, err
}

func (s *sendRepo) CreatEmail(emailTextId string, email string, status bool) error {
	query := `INSERT INTO email_send_email(
			id,
			email_text_id,
			email,
			send_status,
			created_at)
			VALUES($1, $2, $3, $4, $5);`

	fmt.Println("\n\n\n\n>>>>>>>>>>>>>", emailTextId, email, status)

	err := s.db.DB.QueryRow(
		query,
		uuid.New().String(),
		emailTextId,
		email,
		status,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
	).Err()

	if err != nil {
		return err
	}
	return nil
}

func (s *sendRepo) CreatSms(text string, phone string) error {
	query := `INSERT INTO sms_send_sms(
			id,
			text,
			phone,
			created_at)
			VALUES($1, $2, $3, $4);`

	err := s.db.DB.QueryRow(
		query,
		uuid.New().String(),
		text,
		phone,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
	).Err()

	if err != nil {
		return err
	}
	return nil
}
