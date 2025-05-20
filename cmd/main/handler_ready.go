package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	payload := map[string]string{"status": "ok"}
	logrus.Info("Health check endpoint called")
	respondWithJSON(w, http.StatusOK, payload)
}

