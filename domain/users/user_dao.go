package users

import (
	"fmt"

	"github.com/claudiocleberson/bookstore_users-api/datasources/mysql/users_db"
	"github.com/claudiocleberson/bookstore_users-api/logger"
	"github.com/claudiocleberson/bookstore_utils-shared/utils/date_utils"
	"github.com/claudiocleberson/bookstore_utils-shared/utils/mysql_utils"
	"github.com/claudiocleberson/bookstore_utils-shared/utils/rest_err"
)

const (
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryUpdateUser             = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?"
	queryInsertUser             = "INSERT INTO users(first_name, last_name,email,date_created,status,password) VALUES(?,?,?,?,?,?)"
	queryFindUserByStatus       = "SELECT id, first_name,last_name,email,date_created,status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name,last_name,email,date_created,status FROM users WHERE email=? AND password=? AND status=?;"
)

//Mock the database layer
var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *rest_err.RestErr {

	stmt, err := users_db.UsersDB.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare getUser statement", err)
		return rest_err.NewInternalServerError("Database failed")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.Firstname, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while scanning result from getUser method", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByEmailAndPassword() *rest_err.RestErr {

	stmt, err := users_db.UsersDB.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare getUser By email and password statement", err)
		return rest_err.NewInternalServerError("Database failed")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if err := result.Scan(&user.Id, &user.Firstname, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error while scanning result from get user by email and password method", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *rest_err.RestErr {

	//Create the sql statement
	stmt, err := users_db.UsersDB.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error preparing SQL statement for saveUser method", err)
		return rest_err.NewInternalServerError("database error")
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	user.Status = StatusActive

	//Execute quert against DB
	insertResult, saveErr := stmt.Exec(user.Firstname, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error running SQL query for saveUser method", saveErr)
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while  saving new user", err)
		return rest_err.NewInternalServerError("database error: couldn't save user")
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *rest_err.RestErr {

	stmt, err := users_db.UsersDB.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error preparing SQL statement for Update user method", err)
		return rest_err.NewInternalServerError("database error")
	}
	defer stmt.Close()

	valErr := user.Validate()
	if valErr != nil {
		return valErr
	}

	_, err = stmt.Exec(user.Firstname, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error executing SQL statement for Update user method", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *rest_err.RestErr {
	stmt, err := users_db.UsersDB.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error preparing SQL statement for Delete user method", err)
		return rest_err.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error executing SQL statement for Delete user method", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_err.RestErr) {
	stmt, err := users_db.UsersDB.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error preparing SQL statement for FindByStatus user method", err)
		return nil, rest_err.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error running SQL statement FindByStatus user method", err)
		return nil, rest_err.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Firstname, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error scanning rows at FindByStatus user method", err)
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_err.NewNotFoundError(fmt.Sprintf("no user found with the given status %s", status))
	}

	return results, nil

}
