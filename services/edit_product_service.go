package services

import (
	"apigolang/database"
	"apigolang/models"
)

func EditProductService(id string) models.Product {
	db := database.ConnectionDB()
	defer db.Close()

	prod := database.Edit(id)
	return prod
}
