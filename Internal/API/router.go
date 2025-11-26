package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) routes() http.Handler {

	mux := mux.NewRouter()

	mux.HandleFunc("/trans/get/{month,year}", api.getTrans).Methods("GET")
	mux.HandleFunc("/trans/add", api.addTrans).Methods("POST")
	mux.HandleFunc("/trans/del", api.delTrans).Methods("DELETE")

	return mux
}
