package controllers

import (
	"apigolang/services"
	"net/http"
)

func DeleteProductController(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	services.DeleteProductService(id)
	http.Redirect(w, r, "/", 301)
}
