package main

import (
	"apigolang/routes"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	logger.Info("initialized main")

	routes.Routes()

	logger.Info("server is running")
	http.ListenAndServe(":8000", nil)
}
