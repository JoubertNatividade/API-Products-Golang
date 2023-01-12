package controllers

import (
	"apigolang/services"
	"html/template"
	"net/http"
)

var TemplatesPages = template.Must(template.ParseGlob("templates/*.html"))

func ListAllController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		result := services.ListAllService()

		TemplatesPages.ExecuteTemplate(w, "Index", result)
	}
}
