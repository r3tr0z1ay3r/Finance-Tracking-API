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
	DB *sql.DB
}

func NewAPI(dbConn *sql.DB) *API {

	return &API{DB: dbConn}

}

func (api *API) Handle_insertDB(w http.ResponseWriter, r *http.Request) {

	var trans Model.Transaction

	if err := json.NewDecoder(r.Body).Decode(&trans); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	err := Db.InsertDb(api.DB, trans)

	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(trans)

}

func (api *API) Handle_getDB(w http.ResponseWriter, r *http.Request) {

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
	trans, err := Db.GetValDb(api.DB, month, year, userStr)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trans)

}

func (api *API) Handle_delDB(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {

		http.Error(w, "Invalid ID", http.StatusBadRequest)

	}

	err = Db.DeleteValDb(api.DB, id)

	if err != nil {

		http.Error(w, "Error while attempting to delete", http.StatusExpectationFailed)

	}

}
