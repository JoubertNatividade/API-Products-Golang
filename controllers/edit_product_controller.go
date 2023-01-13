package controllers

import (
	"apigolang/services"
	"net/http"
)

func EditProductController(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	prod := services.EditProductService(id)
	TemplatesPages.ExecuteTemplate(w, "Edit", prod)
}
