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

func prepareStatement(sqlCommand string) *sql.Stmt {
	stmt, err := db.Prepare(sqlCommand)
	if err != nil {
		logger.Errorf("can't prepare statement. Here the reason: %s", err)
	}

	return stmt
}

func Insert(
	name,
	description string,
	price float64,
	quantity int,
) {
	stmt := prepareStatement("insert into products (name, descritpion, price, quantity) values ($1,$2,$3, $4)")

	exec, err := stmt.Exec(name, description, price, quantity)

	if err != nil {
		logger.Error("can't insert product")
	} else {
		logger.Info("product insert success")
	}

	exec.RowsAffected()
}
