package Api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) Routes() http.Handler {

	mux := mux.NewRouter()

	//mux.HandleFunc("/trans/get/{month,year}", api.Handle_insertDB).Methods("GET")
	mux.HandleFunc("/trans/add", api.Handle_insertDB).Methods("POST")
	//mux.HandleFunc("/trans/del", api.delTrans).Methods("DELETE")

	return mux
}
