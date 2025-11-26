package Api

import (
	"database/sql"
	"encoding/json"
	"net/http"

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
