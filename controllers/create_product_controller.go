package controllers

import (
	"apigolang/services"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CreateProductController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertPriceInFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			logger.Errorf("can't convert price in float. Here the reason: %s", err)
		}
		convertQuantityInInt, err := strconv.Atoi(quantity)
		if err != nil {
			logger.Errorf("can't convert quantity in int. Here the reason: %s", err)
		}

		services.CrateProductService(name, description, convertPriceInFloat, convertQuantityInInt)
	}

}
