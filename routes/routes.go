package routes

import (
	"apigolang/controllers"
	"net/http"
)

func Routes() {

	http.HandleFunc("/", controllers.ListAllController)
	http.HandleFunc("/new", controllers.NewProductController)
	http.HandleFunc("/store", controllers.CreateProductController)
	http.HandleFunc("/delete", controllers.DeleteProductController)

}
