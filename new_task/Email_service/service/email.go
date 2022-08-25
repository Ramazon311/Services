package service

import (
	"context"
	"log"

	"bytes"
	"encoding/json"
	"github/Services/newpro/Email_service/config"
	l "github/Services/newpro/Email_service/pkg/logger"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github/Services/newpro/Email_service/genproto/email"
	"github/Services/newpro/Email_service/storage"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	gomail "gopkg.in/gomail.v2"
)

// SendService ...
type SendService struct {
	storage storage.I
	Conf    config.Config
	DB      *sqlx.DB
	logger  l.Logger
}

// NewSendService ...
func NewSendService(db *sqlx.DB, conf config.Config) *SendService {
	return &SendService{storage: storage.NewStoragePg(db), Conf: conf}
}

//Send ...
func (s *SendService) Send(ctx context.Context, req *pb.EmailText) (*pb.Empty, error) {

	id, err := s.storage.SendEmail().CreatEmailText(
		uuid.New().String(), req.Subject, req.Body,
	)

	fmt.Println("\n\n\n\n>>>>>>>",id, req.Body, req.Phone, req.Recipints, req.Subject)

	if err != nil {
		s.logger.Error("\n\n\nError creat emailText", l.Error(err))
		return nil, err
	}

	for _, val := range req.Phone {

		if err := s.sendSms(val, req.Body); err != nil {
			s.logger.Error("\n\n\nNo message was sent to"+val, l.Error(err))
		}

	}

	for _, val := range req.Recipints {
		status := true

		if err := s.sendEmail(req.Subject, req.Body, val); err != nil {
			s.logger.Error("\n\n\n>>>>>>>\nNo message was sent to"+val, l.Error(err))
			status = false
		}

		err = s.storage.SendEmail().CreatEmail(
			id,
			val,
			status,
		)
		if err != nil {
			s.logger.Error("\n\n\n>>>>>>\nError creat posgresql"+val, l.Error(err))

		}
	}

	return &pb.Empty{}, nil
}

func (s *SendService) sendEmail(subject, body, email string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", s.Conf.EmailFromHeader)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Send the email to
	d := gomail.NewDialer(s.Conf.SMTPHost, s.Conf.SMTPPort, s.Conf.SMTPUser, s.Conf.SMTPUserPass)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Print("Sent")
	return nil
}

func (s *SendService) sendSms(phone string, text string) error {

	err := s.storage.SendEmail().CreatSms(text, phone)

	if err != nil {
		s.logger.Error("\n\n\n>>>>>>\nError creat posgresql")
	}


	// SMSRequestBody ...
	type SMSRequestBody struct {
		From      string `json:"from"`
		Text      string `json:"text"`
		To        string `json:"to"`
		APIKey    string `json:"api_key"`
		APISecret string `json:"api_secret"`
	}

	body := SMSRequestBody{
		From:      "Coder",
		To:        phone,
		Text:      text,
		APIKey:    "5126ddae",
		APISecret: "TC3LhaIjxTPNsO1p",
	}
	

	smsBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(smsBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respBody))

	return nil
}
