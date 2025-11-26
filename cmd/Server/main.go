package main

import (
	"log"
	"net/http"

	Api "github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/API"
	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/Db"
)

func main() {

	dbConn, err := Db.ConnectDb("./test.db")
	if err != nil {
		log.Fatalf("Error occured while intializing DB Connection : %v\n", err)
	}
	defer dbConn.Close()

	err = Db.CreateDb(dbConn) //Creates table transaction if it does not exists already
	if err != nil {
		log.Fatalf("Error occured while creating a new table : %v\n", err)
	}

	api := Api.NewAPI(dbConn)

	server := &http.Server{
		Addr:    ":8080",
		Handler: api.Routes(),
	}
	log.Println("Server is listening on 8080")
	log.Fatal(server.ListenAndServe())

}
