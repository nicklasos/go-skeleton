package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func dbConnect() error {
	var err error
	db, err = sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		logger.Error("Cannot open connection to mysql")
		return err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Cannot talk to mysql")
		return err
	}

	return nil
}
