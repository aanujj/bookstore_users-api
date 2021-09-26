package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root", "Eg0mania#", "localhost", "users_db",
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		fmt.Println("unable to connect to database")
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
