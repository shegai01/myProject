package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CustomRouter() {
	router := mux.NewRouter()
	log.Println(http.ListenAndServe(":8080", router))
}
