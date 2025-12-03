package Api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) Routes() http.Handler {

	mux := mux.NewRouter()

	protected := mux.PathPrefix("/").Subrouter()
	protected.Use(AuthMiddleWare)

	protected.HandleFunc("/trans/get/{month}/{year}/{user}", api.Handle_getDB).Methods("GET")
	protected.HandleFunc("/trans/add", api.Handle_insertDB).Methods("POST")
	protected.HandleFunc("/trans/del/{id}", api.Handle_delDB).Methods("DELETE")

	return mux
}
