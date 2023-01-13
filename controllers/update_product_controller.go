package controllers

import (
	"apigolang/services"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func UpdateProductController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvertToInt, err := strconv.Atoi(id)
		if err != nil {
			logger.Errorf("can't convertrion id to int. Here the reazon: %s", err)
		}

		priceConvertInFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			logger.Errorf("can't convertrion price to float. Here the reazon: %s", err)
		}

		quantityConvertToInt, err := strconv.Atoi(quantity)
		if err != nil {
			logger.Errorf("can't convertrion quantity to int. Here the reazon: %s", err)
		}

		if (name == "") || (description == "") || (price == "") || (quantity == "") {
			logger.Errorf("can't create producut. Here the reazon: %s", err)
		}

		services.UpdateProductService(idConvertToInt, name, description, priceConvertInFloat, quantityConvertToInt)
	}
	http.Redirect(w, r, "/", 301)
}
