package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type User struct {
	Id            string `json:"id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Profile_photo string `json:"profile_photo"`
	Bio           string `json:"bio"`
	Email         string `json:"Email"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Phone_number  string `json:"phone_number"`
	Password      string `json:"password"`
	Code          string `json:"code"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted_at    string `json:"deleted_at"`
}

type GetProfileByJwtRequestModel struct {
	Token string `json:"token"`
}

type CreateUser struct {
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Profile_photo string `json:"profile_photo"`
	Bio           string `json:"bio"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Phone_number  string `json:"phone_number"`
}

type UpdateUser struct {
	ID            string `json:"id"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Profile_photo string `json:"profile_photo"`
	Bio           string `json:"bio"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Phone_number  string `json:"phone_number"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponseModel struct {
	Message string `json:"message"`
}

type VerifyResponseModel struct {
	ID           string `json:"id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func (rum *User) Validate() error {
	return validation.ValidateStruct(
		rum,
		validation.Field(&rum.First_name, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.Last_name, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.Username, validation.Required, validation.Length(5, 30), validation.Match(regexp.MustCompile("^[0-9a-z_.]+$"))),
		validation.Field(&rum.Email, validation.Required, is.Email),
		validation.Field(&rum.Password, validation.Required, validation.Length(8, 30), Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
	)
}
