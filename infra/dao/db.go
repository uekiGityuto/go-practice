package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "member:@(127.0.0.1:3306)/practice")
	if err != nil {
		panic(err)
	}
	return db
}
