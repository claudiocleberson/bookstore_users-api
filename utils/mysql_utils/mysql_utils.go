package mysql_utils

import (
	"fmt"

	"github.com/claudiocleberson/bookstore_users-api/utils/rest_err"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *rest_err.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		return rest_err.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s ", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return rest_err.NewBadRequestError("email already exists")
	}

	return rest_err.NewInternalServerError(fmt.Sprintf("error processing request %s", sqlErr.Message))

}
