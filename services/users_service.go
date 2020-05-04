package services

import (
	"federicoleon/bookstore_oauth-api/src/utils/crypto_utils"

	"github.com/claudiocleberson/bookstore_users-api/domain/users"
	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(user users.User) (*users.User, *rest_err.RestErr)
	GetUser(userId int64) (*users.User, *rest_err.RestErr)
	UserLogin(logiRequest users.LoginRequest) (*users.User, *rest_err.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *rest_err.RestErr)
	DeleteUser(userId int64) *rest_err.RestErr
	FindByStatus(status string) (users.Users, *rest_err.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, *rest_err.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *rest_err.RestErr) {

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

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_err.RestErr) {

	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.Email = user.Email
		current.Firstname = user.Firstname
		current.LastName = user.LastName
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *rest_err.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()

}

func (s *usersService) FindByStatus(status string) (users.Users, *rest_err.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) UserLogin(loginRequest users.LoginRequest) (*users.User, *rest_err.RestErr) {

	dao := &users.User{Email: loginRequest.Email, Password: crypto_utils.GetMd5(loginRequest.Password)}

	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
