package routes

import (
	"apigolang/controllers"
	"html/template"
	"net/http"
)

var TemplatesPages = template.Must(template.ParseGlob("templates/*.html"))

func Routes() {

	http.HandleFunc("/new", controllers.CreateProductController)

}
