package services

import "apigolang/database"

func UpdateProductService(
	id int,
	name,
	description string,
	price float64,
	quantity int,
) {

	db := database.ConnectionDB()
	defer db.Close()

	database.Update(id, name, description, price, quantity)
}
