package main

import (
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	logger.Info("initialized main")

	logger.Info("server is running")
	http.ListenAndServe(":8000", nil)
}
