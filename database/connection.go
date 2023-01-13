package database

import (
	"apigolang/models"
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
func ListAll() *sql.Rows {

	all, _ := db.Query("select * from products order by id asc")

	return all
}

func Insert(
	name,
	description string,
	price float64,
	quantity int,
) {
	stmt := prepareStatement("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")

	exec, err := stmt.Exec(name, description, price, quantity)

	if err != nil {
		logger.Error("can't insert product")
	} else {
		logger.Info("product insert success")
	}

	exec.RowsAffected()
}

func Edit(id string) models.Product {
	response, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		logger.Errorf("can't select produtc. Here the reason: %s", err)
	}

	prod := models.Product{}

	for response.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := response.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			logger.Errorf("can't finc item in db. Here the reason: %s", err)
		}

		prod.ID = id
		prod.Name = name
		prod.Description = description
		prod.Price = price
		prod.Quantity = quantity
	}
	return prod
}

func Update(
	id int,
	name,
	description string,
	price float64,
	quantity int,
) {
	stmt := prepareStatement("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	stmt.Exec(name, description, price, quantity, id)

}

func Delete(id string) {
	stmt := prepareStatement("delete from products where id=$1")

	stmt.Exec(id)
}
