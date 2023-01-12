package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

var (
	db  *sql.DB
	err error
)

func ConnectionDB() *sql.DB {
	connStr := "user=postgres dbname=postgres password=1234 host=localhost  sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		logger.Errorf("connection database refused. Here the reason:%s,", err)
	}
	return db
}
