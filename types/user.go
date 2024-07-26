package types

import (
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	COST           = 12
	MIN_FIRST_NAME = 2
	MIN_LAST_NAME  = 2
	MIN_PASSWORD   = 7
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateUsersParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Pass      string `json:"password"`
}

func (u UpdateUserParams) ToBSON() bson.M {
	d := bson.M{}
	if len(u.FirstName) > 0 {
		d["firstName"] = u.FirstName
	}
	if len(u.LastName) > 0 {
		d["lastName"] = u.LastName
	}
	return d
}

func (c *CreateUsersParams) Validate() map[string]string {
	errors := make(map[string]string)
	if len(c.FirstName) < MIN_FIRST_NAME {
		errors["firstName"] = fmt.Sprintf("first name should be at least %d characters", MIN_FIRST_NAME)
	}
	if len(c.LastName) < MIN_LAST_NAME {
		errors["lastName"] = fmt.Sprintf("last name should be at least %d characters", MIN_LAST_NAME)
	}
	if len(c.Pass) < MIN_PASSWORD {
		errors["password"] = fmt.Sprintf("password should be at least %d characters", MIN_PASSWORD)
	}

	if !isEmailValid(c.Email) {
		errors["email"] = "invalid email"
	}

	return errors
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

type User struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName     string             `json:"firstName" bson:"firstName"`
	LastName      string             `json:"lastName" bson:"lastName"`
	Email         string             `json:"email" bson:"email"`
	EncryptedPass string             `json:"-" bson:"encryptedPass"`
}

func NewUserFromParams(params CreateUsersParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Pass), COST)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:     params.FirstName,
		LastName:      params.LastName,
		Email:         params.Email,
		EncryptedPass: string(encpw),
	}, nil

}
