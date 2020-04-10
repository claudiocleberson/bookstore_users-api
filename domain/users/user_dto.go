package users

import (
	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *rest_err.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower((user.Email)))
	if user.Email == "" {
		return rest_err.NewBadRequestError("invalid email addresss")
	}
	return nil
}
