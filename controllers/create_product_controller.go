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

		priceConvertInFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			logger.Errorf("can't convertrion value to float. Here the reazon: %s", err)
		}

		quantityConvertToInt, err := strconv.Atoi(quantity)
		if err != nil {
			logger.Errorf("can't convertrion value to int. Here the reazon: %s", err)
		}

		if (name == "") || (description == "") || (price == "") || (quantity == "") {
			logger.Errorf("can't create producut. Here the reazon: %s", err)
		} else {
			services.CrateProductService(name, description, priceConvertInFloat, quantityConvertToInt)
		}
	}
	http.Redirect(w, r, "/", 301)
}
