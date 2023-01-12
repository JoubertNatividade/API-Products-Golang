package controllers

import "net/http"

func NewProductController(w http.ResponseWriter, r *http.Request) {
	TemplatesPages.ExecuteTemplate(w, "New", nil)
}
