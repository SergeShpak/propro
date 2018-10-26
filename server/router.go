package main

import (
	"github.com/gorilla/mux"

	"github.com/SergeyShpak/propro/server/handlers"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/ping").Methods("GET").HandlerFunc(handlers.Ping)
	return r
}
