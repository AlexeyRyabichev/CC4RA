package public_api

import (
	"github.com/gorilla/mux"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/internal/matrix"
	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

func GetHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/matrix", matrix.CountMatrices).
		Methods("POST")

	router.Use(logger.LoggingMiddleware)

	return router
}
