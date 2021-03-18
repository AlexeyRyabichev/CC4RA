package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

type API struct {
	Server        *http.Server
	Router        *mux.Router
	Specification *Specification
}

func NewAPI(spec *Specification, router *mux.Router) (*API, error) {
	server := &http.Server{
		Addr:         spec.GetAddr(),
		WriteTimeout: spec.WriteTimeout,
		ReadTimeout:  spec.ReadTimeout,
		IdleTimeout:  spec.IdleTimeout,
		Handler:      router,
	}

	return &API{Server: server, Router: router, Specification: spec}, nil
}

func (api *API) Run() {
	if err := api.Server.ListenAndServe(); err != nil {
		logger.Sugar.Error(err.Error())
	}
}
