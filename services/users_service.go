package services

import (
	"github.com/claudiocleberson/bookstore_users-api/domain/users"
	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
)

func CreateUser(user users.User) (*users.User, *rest_err.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *rest_err.RestErr) {

	if userId <= 0 {
		return nil, rest_err.NewBadRequestError("id cannot be 0 or negative")
	}

	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}
