package services

import "apigolang/database"

func DeleteProductService(id string) {

	db := database.ConnectionDB()
	defer db.Close()

	database.Delete(id)

}
