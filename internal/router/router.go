package router

import (
	"github.com/gorilla/mux"
)

func CustomRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}
