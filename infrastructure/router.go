package infrastructure

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func ListenServe(router *mux.Router) {
	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
