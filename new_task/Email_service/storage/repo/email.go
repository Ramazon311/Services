package repo

// Email ...
type Email struct {
	ID             string
	Subject        string
	Body           string
	RecipientEmail string
}

// SendStorageI ...
type SendStorageI interface {
	CreatEmailText(id, subject, body string) (string, error)
	CreatEmail(emailTextId string, email string,  status bool) error
	CreatSms(text string, phone string) error
}