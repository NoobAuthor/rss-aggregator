package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	logrus.Error("HandlerError endpoint was called")
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
