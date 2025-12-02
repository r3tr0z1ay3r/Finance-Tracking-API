//go:build ignore

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Import SQLite3 driver
)

type model struct {
	ID   int       `json:"id"`
	Time time.Time `json:"time"`
	Amt  float64   `json:"amt"`
	Flow string    `json:"flow"`
	Mode string    `json:"mode"`
	User string    `json:"usr_name"`
}

func connectDb() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {

		return nil, err

	}

	if err = db.Ping(); err != nil {

		return nil, err

	}
	return db, err

}

func createDb(db *sql.DB) error {

	// Function to create a table called transactions
	// Expects Database connection to be provided as input and returns any err if occured
	// Expected usecase : first run

	//Create a table
	createTableTrans := `CREATE TABLE IF NOT EXISTS transactions (
						id	INTEGER PRIMARY KEY AUTOINCREMENT,
						time DATETIME NOT NULL,
						amt REAL NOT NULL,
						flow TEXT NOT NULL,
						mode TEXT,
						user TEXT
	);`

	createTableUser := `CREATE TABLE IF NOT EXISTS users (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						name TEXT NOT NULL,
						pass VARCHAR 
	);`

	_, err := db.Exec(createTableTrans)
	if err != nil {
		return err
	}

	_, err = db.Exec(createTableUser)
	if err != nil {
		return err
	}

	return nil

}

func insertDb(db *sql.DB, amt float64, mode string, flow string, user string) error {

	// Function to insert value onto the database
	// Expects database, (Amount , Mode of payment, Expense/Income) -> Should be parsed from API
	// Should return error if any else Adds value to database
	currTime := time.Now()
	formattedTime := currTime.Format("2006-01-02 15:04:04")
	//time, _ := time.Parse("2006-01-02 15:04:04", formattedTime)
	fmt.Printf("Current Time is %v\n", formattedTime)
	fmt.Printf("The values from the api are %v,%v,%v, %v\n", amt, flow, mode, user)
	insertCmd := `INSERT INTO transactions(time, amt, flow, mode, user) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertCmd, formattedTime, amt, flow, mode, user)
	return err
}

func getValDb(db *sql.DB, month int, year int, user string) {
	// Function to fetch values from the database
	// Designed to get values based on the month and year
	// Designed into months and years to get monthly transactions on a page

	yr_mth := fmt.Sprintf("%v-%v", year, month)
	cmd := fmt.Sprintf(`
		SELECT *
		FROM transactions
		WHERE user = '%v' AND strftime('%%Y-%%m',time) = '%v'`, user, yr_mth)

	val, err := db.Query(cmd)
	if err != nil {

		log.Fatalf("Error while fetching rows : %v ", err)

	}
	defer val.Close()

	var Rows []model

	for val.Next() {
		var t model
		err := val.Scan(&t.ID, &t.Time, &t.Amt, &t.Flow, &t.Mode, &t.User)
		if err != nil {
			log.Fatal(err)
		}
		Rows = append(Rows, t)

	}

	fmt.Printf("%v\n", Rows)

}

func printVals(rows *sql.Rows) {

	for rows.Next() {

		var id int
		var time time.Time
		var amt float64
		var flow string
		var mode string

		err := rows.Scan(&id, &time, &amt, &flow, &mode)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v |\n", id, time, amt, flow, mode)
	}

}

func deleteValDb(db *sql.DB, id int) error {

	delCmd := fmt.Sprintf(`DELETE FROM transactions WHERE id = %d`, id)
	_, err := db.Exec(delCmd)
	return err

}
func main() {

	db, err := connectDb()
	if err != nil {

		log.Fatalf("Error in connecting to DB :%v\n", err)

	}
	defer db.Close()

	if err = createDb(db); err != nil {

		log.Fatalf("Error while creating db: %v\n", err)

	}
	// if err = insertDb(db, 67, "UPI", "Expense", "dummy1"); err != nil {

	// 	log.Fatalf("Error while inserting : %v\n", err)

	// }

	deleteValDb(db, 16)

	if err = deleteValDb(db, 8); err != nil {
		log.Fatal(err)
	}

}
