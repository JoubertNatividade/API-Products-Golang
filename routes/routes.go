package routes

import (
	"apigolang/controllers"
	"net/http"
)

func Routes() {

	http.HandleFunc("/", controllers.ListAllController)
	http.HandleFunc("/new", controllers.CreateProductController)

}
