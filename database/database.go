package database

import (
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB = Connection()

func Connection() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=postgres password=123456789 dbname=test sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
