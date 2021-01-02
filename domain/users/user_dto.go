// dto is object that you would transfer from the db to the application and back

package users

import (
	"github.com/davidlamyc/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
	Password string `json:"password"`
}

type Users []User

// Using a method perferred
// func Validate(user *User) *errors.RestErr {
//	 user.Email = strings.TrimSpace(strings.ToLower(user.Email))
//	 if user.Email == "" {
//		 return errors.NewBadRequestError("invalid email address")
//	 }
//	 return nil
// }

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}