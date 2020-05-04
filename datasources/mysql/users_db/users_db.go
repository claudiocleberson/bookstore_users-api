package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	UsersDB *sql.DB
)

const (
	mysql_user_username = "root"
	mysql_user_password = "msb1679"
	mysql_user_host     = "127.0.0.1:3306"
	mysql_user_schema   = "users_db"
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysql_user_username,
		mysql_user_password,
		mysql_user_host,
		mysql_user_schema,
	)

	var err error
	UsersDB, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = UsersDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")

}
