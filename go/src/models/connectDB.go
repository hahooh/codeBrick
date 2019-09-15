package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "jw:jiwankim@tcp(db:3306)/test_db")

	if err != nil {
		panic(err.Error())
	}

	return db
}
