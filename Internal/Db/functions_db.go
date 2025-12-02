package Db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Internal/Model"
	_ "modernc.org/sqlite" // Import SQLite3 driver
)

func CreateDb(db *sql.DB) error {
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
						name TEXT NOT NULL
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

func InsertDb(db *sql.DB, m Model.Transaction) error {

	// Function to insert value onto the database
	// Expects database, (Amount , Mode of payment, Expense/Income) -> Should be parsed from API
	// Should return error if any else Adds value to database
	amt := m.Amt
	flow := m.Flow
	mode := m.Mode
	currTime := time.Now()
	formattedTime := currTime.Format("2006-01-02 15:04:04")
	m.Time, _ = time.Parse("2006-01-02 15:04:04", formattedTime)
	fmt.Printf("Current Time is %v\n", formattedTime)
	fmt.Printf("The values from the api are %v,%v,%v\n", amt, flow, mode)
	insertCmd := `INSERT INTO transactions(time, amt, flow, mode, user) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertCmd, formattedTime, amt, flow, mode)
	return err
}

func GetValDb(db *sql.DB, month int, year int, user string) ([]Model.Transaction, error) {

	// Function to fetch values from the database
	// Designed to get values based on the month and year
	// Designed into months and years to get monthly transactions on a page

	//Creating a map to covert month in english to integer {String(Month) -> Int(month)}

	yr_mth := fmt.Sprintf("%v-%v", year, month)
	cmd := fmt.Sprintf(`
		SELECT *
		FROM transactions
		WHERE strftime('%%Y-%%m',time) = '%v' AND user = %v`, yr_mth, user)

	val, err := db.Query(cmd)
	if err != nil {

		log.Fatalf("Error while fetching rows : %v ", err)

	}
	defer val.Close()

	var Rows []Model.Transaction

	for val.Next() {

		var t Model.Transaction

		err := val.Scan(&t.ID, &t.Time, &t.Amt, &t.Flow, &t.Mode)
		if err != nil {
			log.Fatal(err)
		}
		Rows = append(Rows, t)

	}

	fmt.Printf("%v\n", Rows)
	return Rows, nil

}

func PrintVals(rows *sql.Rows) {

	//Helper function to print the rows gathered from getValDb()
	//Used while debugging
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
		fmt.Printf("%v |%v |%v |%v |%v |\n", id, time, amt, flow, mode)
	}

}

func DeleteValDb(db *sql.DB, id int) error {

	// Function to delete values from the database
	// Expects the database connection (db) and id to delete from
	// Planning to delete val from database based on selection on the front-end,
	// Front-End should only show necessary info, ID will be abstracted but used here

	delCmd := fmt.Sprintf(`DELETE FROM transactions WHERE id = %d`, id)
	_, err := db.Exec(delCmd)
	return err

}
