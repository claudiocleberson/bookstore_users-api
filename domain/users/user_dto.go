package users

import (
	"strings"

	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *rest_err.RestErr {
	user.Firstname = strings.TrimSpace(user.Firstname)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower((user.Email)))
	if user.Email == "" {
		return rest_err.NewBadRequestError("invalid email addresss")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_err.NewBadRequestError("the password cannot be empty")
	}
	return nil
}
