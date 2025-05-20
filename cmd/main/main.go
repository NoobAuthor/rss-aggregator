package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logrus.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.WithError(err).Warn(".env file not found or failed to load")
	}

	portString := getEnvOrFatal("PORT")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1ROUTER := chi.NewRouter()
	v1ROUTER.Get("/health", HandlerReadiness)
	v1ROUTER.Get("/error", HandlerError)

	r.Mount("/v1", v1ROUTER)

	srv := &http.Server{
		Handler: r, Addr: ":" + portString,
	}

	logrus.WithFields(logrus.Fields{
		"port": portString,
	}).Info("Starting server")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("Server failed to start or stopped unexpectedly")
	}
}
