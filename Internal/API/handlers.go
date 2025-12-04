package Api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/Db"
	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/Model"
)

type API struct {
	Trans_DB *sql.DB
	User_DB  *sql.DB
}

func NewAPI(dbConn_trans *sql.DB, dbConn_user *sql.DB) *API {

	return &API{Trans_DB: dbConn_trans, User_DB: dbConn_user}

}

func (api *API) Handle_insertDBTrans(w http.ResponseWriter, r *http.Request) {

	var trans Model.Transaction

	if err := json.NewDecoder(r.Body).Decode(&trans); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	err := Db.InsertDbTrans(api.Trans_DB, trans)

	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(trans)

}

func (api *API) Handle_getDBTrans(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	monthStr := vars["month"]
	yearStr := vars["year"]
	userStr := vars["user"]
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		http.Error(w, "Invalid Month", http.StatusBadRequest)
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
	}
	trans, err := Db.GetValDbTrans(api.Trans_DB, month, year, userStr)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trans)

}

func (api *API) Handle_delDBTrans(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {

		http.Error(w, "Invalid ID", http.StatusBadRequest)

	}

	err = Db.DeleteValDbTrans(api.Trans_DB, id)

	if err != nil {

		http.Error(w, "Error while attempting to delete", http.StatusExpectationFailed)

	}

}
