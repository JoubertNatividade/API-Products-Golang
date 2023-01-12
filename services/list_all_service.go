package services

import (
	"apigolang/database"
	"apigolang/models"
	logger "github.com/sirupsen/logrus"
)

func ListAllService() []models.Product {
	db := database.ConnectionDB()
	defer db.Close()

	selectAll := database.ListAll()

	prod := models.Product{}
	var products []models.Product

	for selectAll.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAll.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			logger.Errorf("can't scanner items. Here the reason: %s", err)
		}

		prod.ID = id
		prod.Name = name
		prod.Description = description
		prod.Price = price
		prod.Quantity = quantity

		products = append(products, prod)
	}

	defer db.Close()

	return products
}
