package services

import "apigolang/database"

func CrateProductService(
	name,
	description string,
	price float64,
	quantity int,
) {

	db := database.ConnectionDB()
	defer db.Close()

	database.Insert(name, description, price, quantity)

}
