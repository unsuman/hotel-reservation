package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	COST = 12
)

type CreateUsersParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Pass      string `json:"password"`
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
