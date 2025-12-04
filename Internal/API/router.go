package Api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) Routes() http.Handler {

	mux := mux.NewRouter()

	protected := mux.PathPrefix("/").Subrouter()
	protected.Use(AuthMiddleWare)

	protected.HandleFunc("/trans/get/{month}/{year}/{user}", api.Handle_getDBTrans).Methods("GET")
	protected.HandleFunc("/trans/add", api.Handle_insertDBTrans).Methods("POST")
	protected.HandleFunc("/trans/del/{id}", api.Handle_delDBTrans).Methods("DELETE")

	// protected.HandleFunc("/user/login/{user}/{pass_hash}", api.Handle_LoginUser).Methods("GET")
	// protected.HandleFunc("/user/signup", api.Handle_SignupUser).Methods("POST")
	// protected.HandleFunc("/user/remove/{user}/{pass_hash}", api.Handle_RemoveUser).Methods("DELETE")

	return mux
}
