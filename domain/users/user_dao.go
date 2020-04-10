package users

import (
	"fmt"

	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
)

//Mock the database layer
var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *rest_err.RestErr {

	result := usersDB[user.Id]
	if result == nil {
		return rest_err.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Firstname = result.Firstname
	user.LastName = result.LastName
	user.Email = result.Email

	return nil
}

func (user *User) Save() *rest_err.RestErr {

	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return rest_err.NewBadRequestError(fmt.Sprintf("email %v already registered", user.Email))
		}
		return rest_err.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
